package main

import "net"

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

}
