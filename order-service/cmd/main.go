package main

import (
	"L0/order-service/internal/adapters/api"
	"L0/order-service/internal/domain/cache"
	"L0/order-service/internal/domain/config"
	"L0/order-service/internal/domain/nats"
	"L0/order-service/internal/domain/postgres"
)

func main() {

	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := postgres.InitDB(); err != nil {
		panic(err)
	}

	if err := cache.InitCache(); err != nil {
		panic(err)
	}

	go func() {
		if err := nats.InitListener(); err != nil {
			panic(err)
		}
	}()

	if err := api.InitRouter(); err != nil {
		panic(err)
	}

}
