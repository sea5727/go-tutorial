package main

import (
	"fmt"
	"log"
	"net"
	// "golang.org/x/sys/unix"
)


var epoller *epoll

func main() {
	fmt.Println("Hello World")

	ln, err := net.Listen("tcp", ":12345")

	if err != nil {
		panic(err)
	}

	epoller, err = MkEpoll()
	if err != nil {
		panic(err)
	}

	go start()

	for {
		_, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err:%v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
		}

		epoller.

	}
}

func start() {

}
