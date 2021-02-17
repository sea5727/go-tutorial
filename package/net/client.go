package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	go func() {
		data := make([]byte, 8096)

		for {
			len, err := conn.Read(data)
			if err != nil {
				panic(err)
			}

			log.Println("[RECV] : ", string(data[:len]))
		}

	}()

	for {
		time.Sleep(time.Duration(1) * time.Second)
	}

}
