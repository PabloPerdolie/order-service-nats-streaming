package main

import (
	"L0/order-service/internal/adapters/api"
	"L0/order-service/internal/domain/config"
	"L0/order-service/internal/domain/nats"
	"L0/order-service/internal/domain/postgres"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		panic(err)
	}

	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := postgres.InitDB(); err != nil {
		panic(err)
	}

	if err := nats.InitListener(); err != nil {
		panic(err)
	}

	if err := api.InitRouter(); err != nil {
		panic(err)
	}
}
