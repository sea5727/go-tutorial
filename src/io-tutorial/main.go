package main

import (
	"fmt"
	_ "fmt"
	"io/ioutil"
	"os"
	_ "os"
)

func main() {
	data, err := ioutil.ReadFile("./testfile")
	if err != nil {
		panic(err.Error)
	}
	fmt.Println("data:", data)

	write_err := ioutil.WriteFile("./testfilecopy", data, os.FileMode(644))
	if write_err != nil {
		panic(write_err.Error)
	}

	println("Success Bye")
}
