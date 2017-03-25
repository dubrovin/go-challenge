package main

import (
	"time"
	"github.com/dubrovin/go-challnge/server"
	"flag"
	"github.com/prometheus/common/log"
)

func main() {
	timeout := flag.String("timeout", "500ms", "timeout for client")
	listenAddr := flag.String("http.addr", ":8080", "http listen address")
	flag.Parse()

	expectedTimeout, err := time.ParseDuration(*timeout)
	if err != nil {
		log.Fatal("Error in parse duration", err)
	}
	newServer := server.NewServer(*listenAddr, expectedTimeout)
	newServer.Run()
}