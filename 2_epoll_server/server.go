package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"golang.org/x/sys/unix"
	// "golang.org/x/sys/unix"
)

type epoll struct {
	fd          int
	connections map[int]net.Conn
	lock        *sync.RWMutex
}

func main() {
	fmt.Println("Hello World")

	ln, err := net.Listen("tcp", ":12345")

	if err != nil {
		panic(err)
	}

	_, err = MkEpoll()
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

	}
}

func MkEpoll() (*epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}

	return &epoll{
		fd:          fd,
		lock:        &sync.RWMutex{},
		connections: make(map[int]net.Conn),
	}, nil

}

func start() {

}
