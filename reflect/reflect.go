package main

import (
	"fmt"
	"net"
	"reflect"
)

func typeSwitch(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is int.")
	case string:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is string.")
	default:
		fmt.Println("type:", reflect.ValueOf(value).Type())
	}
}

func main() {

	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
	}

	conn, err := listener.Accept()

	valueof := reflect.ValueOf(conn)
	indirect := reflect.ValueOf(conn)
	field := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fmt.Println("valueof:", valueof)
	fmt.Println("indirect:", indirect)
	fmt.Println("field:", field)

	fmt.Println(reflect.ValueOf(conn))
	fmt.Println(reflect.Indirect(reflect.ValueOf(conn)))
	fmt.Println(reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn"))

	fmt.Println("reflect.go")
	typeSwitch(42)
	typeSwitch("42")
	typeSwitch([]float32{3.14})
}
