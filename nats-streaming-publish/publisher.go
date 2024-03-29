package main

import (
	"bufio"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func main() {
	//connect, err := stan.NatsConnect("test-cluster", "pub", stan.NatsURL("nats://localhost:4222"))
	connect, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("failed to connect to nats: %v", err)
		return
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
		return
	}

	js, err := connect.JetStream()
	if err != nil {
		panic(err)
	}

	stream, err := js.StreamInfo("TEST")
	if err != nil {
		log.Println(err)
	}

	if stream == nil {
		log.Println("creating stream and subjects")
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     "TEST",
			Subjects: []string{"TEST.*"},
		})

		if err != nil {
			panic(err)
		}
	}

	err = connect.Publish("TEST.orders", mes)
	if err != nil {
		log.Fatalf("failed to publish: %v", err)
		return
	} else {
		log.Println("mes published")
	}
}
