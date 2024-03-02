package nats

import (
	"L0/order-service/internal/domain/models"
	"L0/order-service/internal/domain/postgres"
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

func InitListener() error {

	natsCon, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("failed to connect nats: %v", err)
		return err
	}

	_, err = natsCon.Subscribe("test.orders", func(msg *nats.Msg) {
		log.Printf("mes received %v", msg.Data)
		var order models.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			return
		}
		postgres.CreateOrder(context.Background(), order)
		//todo CACHE
	})
	if err != nil {
		return err
	}
	return nil
}
