package nats

import (
	"L0/order-service/internal/domain/config"
	"L0/order-service/internal/domain/models"
	"L0/order-service/internal/domain/postgres"
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

func InitListener() error {
	conf := config.CONFIG.Nats
	natsCon, err := nats.Connect(conf.URL)
	if err != nil {
		return err
	}

	js, err := natsCon.JetStream()

	if err != nil {
		return err
	}

	log.Println("Successfully initialized nats-listener")

	_, err = js.Subscribe(conf.SubjectName, func(msg *nats.Msg) {
		msg.Ack()

		var order models.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			return
		}
		log.Printf("mes received %v", order)
		postgres.CreateOrder(context.Background(), order)
		//todo CACHE
	}, nats.Durable(conf.Subscriber), nats.ManualAck())

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}
