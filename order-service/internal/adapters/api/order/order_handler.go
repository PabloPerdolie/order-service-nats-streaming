package order

import (
	"L0/order-service/internal/domain/cache"
	"L0/order-service/internal/domain/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&order); err != nil {
		http.Error(w, "Failed to decode data", http.StatusBadRequest)
		log.Printf("Failed to decode data: %v", err)
		return
	}

	order, err := cache.GetFromCache(order.OrderUid)
	if err != nil {
		http.Error(w, "Failed to find by id", http.StatusBadRequest)
		log.Printf("Failed to find by id: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(order.JsonData)
}

func GetAll(w http.ResponseWriter, r *http.Request) {

}
