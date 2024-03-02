package main

import (
	"L0/order-service/internal/adapters/api"
	"L0/order-service/internal/domain/nats"
	"L0/order-service/internal/domain/postgres"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		panic(err)
	}

	router := mux.NewRouter()

	http.Handle("/", router)
	api.SetupRoutes(router)
	//go func() {
	//	db, err := postgres.InitClient()
	//	if err != nil {
	//		log.Fatalf("failed %v", err)
	//	}
	//	err = db.CreateOrder(context.Background(), models.Delivery{
	//		Name:    "",
	//		Phone:   "",
	//		Zip:     "",
	//		City:    "",
	//		Address: "",
	//		Region:  "",
	//		Email:   "",
	//	})
	//	if err != nil {
	//		log.Fatalf("%v", err)
	//		return
	//	}
	//	log.Printf("success %v", err)
	//}()
	//todo
	db, err := postgres.InitClient()
	if err != nil {
		log.Fatalf("failed %v", err)
	}
	
	go func() {
		err = nats.ListenerConnect(db)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
