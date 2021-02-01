package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	fmt.Println("HelloWorld")

	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
	}

	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp error:%v", ne)
				continue
			}

			log.Printf("accept error:%v", err)
			return
		}

		go handleConn(conn)
		connections = append(connections, conn)
		if len(connections)%100 == 0 {
			log.Printf("total number of connectoins : %v", len(connections))
		}

	}

}

func handleConn(conn net.Conn) {
	io.Copy(ioutil.Discard, conn)
}
