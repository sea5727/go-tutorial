package main

import (
	"fmt"
	"net"
	"sync"
)

// MyStruct for test
type MyStruct struct {
	a int
	b string
	c map[int]string
}

// epoll struct for test
type epoll struct {
	fd          int
	connections map[int]net.Conn
	lock        *sync.RWMutex
}

func main() {
	fmt.Println("Hello World")

	st := MyStruct{}
	st.a = 100

	st2 := MyStruct{100, "test", make(map[int]string)}
	st2.c[0] = "my0"
	fmt.Println("MyStruct : ", st)
	fmt.Println("MyStruct : ", st2)

	st3 := epoll{1000, make(map[int]net.Conn), nil}
	st4 := epoll{}
	fmt.Println("MyStruct : ", st3)
	fmt.Println("MyStruct : ", st4)

}
