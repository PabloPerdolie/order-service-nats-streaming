package main

import (
	"bufio"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	connect, err := stan.Connect("test-cluster", "pub", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("failed to connect to nats: %v", err)
	}
	defer connect.Close()

	open, err := os.Open("nats-streaming-publish/model.json")
	if err != nil {
		log.Fatalf("failed to open model.json: %v", err)
	}

	var mes []byte
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		mes = append(mes, scanner.Bytes()...)
	}
	if scanner.Err() != nil {
		log.Fatalf("failed to scan file: %v", scanner.Err().Error())
	}

	err = connect.Publish("test", mes)
	if err != nil {
		log.Fatalf("failed to publish: %v", err)
	} else {
		log.Println("mes published")
	}
}
