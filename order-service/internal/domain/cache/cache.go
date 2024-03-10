package cache

import (
	"L0/order-service/internal/domain/models"
	"L0/order-service/internal/domain/postgres"
	"context"
	"log"
)

var (
	CACHE map[string]models.Order
)

func InitCache() error {
	CACHE = make(map[string]models.Order)
	orders, err := postgres.GetAll(context.Background())
	if err != nil {
		return err
	}
	for _, order := range orders {
		CACHE[order.OrderUid] = order
	}

	log.Println("Successfully initialized cache")

	return nil
}

func AddToCache(order models.Order) {
	if _, ok := CACHE[order.OrderUid]; ok {
		log.Println("This order is already in cache memory")
		return
	}
	CACHE[order.OrderUid] = order
	log.Println("Successfully added order to cache memory")
}

func GetFromCache(uid string) (order models.Order, err error) {
	order, ok := CACHE[uid]
	if ok {
		log.Printf("From cache with id: %s", uid)
		return order, nil
	}
	order, err = postgres.GetOrderByUid(context.Background(), uid)
	if err != nil {
		return order, err
	}
	log.Printf("From db with id: %s", uid)

	err = InitCache()
	if err != nil {
		return order, err
	}

	return order, nil
}
