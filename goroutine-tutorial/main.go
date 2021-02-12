package main

import (
	"fmt"
	"time"
)

func func1() {
	time.Sleep(time.Second * 1)
	fmt.Println("go func1 end ( Sleep 1 seconds) ")
}

func func2() {
	time.Sleep(time.Second * 2)
	fmt.Println("go func2 end ( Sleep 1 seconds) ")
}

func func3() {
	time.Sleep(time.Second * 3)
	fmt.Println("go func3 end ( Sleep 1 seconds) ")
}
func main() {

	go func1()
	go func2()
	go func3()

	time.Sleep(time.Second * 5)
}
