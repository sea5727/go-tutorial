package main

import (
	"fmt"
	"io"

	"log"
	"os"
)

var myLogger *log.Logger
var fileLogger *log.Logger

func main() {
	fmt.Println("Hello fmt")
	log.Println("Hello log")

	myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

	run()

	myLogger.Println("End of Program")

	filelog()

	multilog()

}

func run() {
	myLogger.Println("Run")
}

func filelog() {
	file, err := os.OpenFile("./program.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	fileLogger = log.New(file, "FILE: ", log.Ldate|log.Ldate|log.Lshortfile)

	fileLogger.Println("File Test")

	defer file.Close()

}

func multilog() {
	file, err := os.OpenFile("./multilog.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := io.MultiWriter(file, os.Stdout)
	log.SetOutput(writer)

	log.Println("MultiFileLog Test")

}
