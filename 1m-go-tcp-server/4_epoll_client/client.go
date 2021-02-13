package main

import (
	"flag"
	"log"
	"time"

	"github.com/rcrowley/go-metrics"
)

var (
	ip          = flag.String("ip", "127.0.0.1", "server IP")
	connections = flag.Int("conn", 1, "number of tcp connections")
	startMetric = flag.String("sm", time.Now().Format("2006-01-02T15:04:05 -0700"), "start time point of all clients")
)

var (
	opsRate = metrics.NewRegisteredTimer("ops", nil)
)

var epoller *epoll

func main() {
	flag.Parse()

	log.Println("HelloWorld!!")

}
