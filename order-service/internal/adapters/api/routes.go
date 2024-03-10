package api

import (
	"L0/order-service/internal/adapters/api/order"
	"L0/order-service/internal/domain/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouter() error {
	router := mux.NewRouter()
	http.Handle("/", router)
	setupRoutes(router)
	log.Println("Server is running: http://localhost:8080/")
	err := http.ListenAndServe(config.CONFIG.Server.URL, nil)
	if err != nil {
		return err
	}
	return nil
}

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/get", order.GetOrder).Methods("GET")
	r.HandleFunc("/getall", order.GetAll).Methods("GET")
}
