package cache

import (
	"L0/order-service/internal/domain/models"
	"L0/order-service/internal/domain/postgres"
	"context"
	"log"
	"sync"
)

var (
	CACHE *sync.Map
)

func InitCache() error {
	CACHE = &sync.Map{}
	orders, err := postgres.GetAll(context.Background())
	if err != nil {
		return err
	}
	for _, order := range orders {
		CACHE.Store(order.OrderUid, order)
	}

	log.Println("Successfully initialized cache")

	return nil
}

func AddToCache(order models.Order) {
	CACHE.Store(order.OrderUid, order)
}

func GetFromCache(uid string) (any, error) {
	order, ok := CACHE.Load(uid)
	if ok {
		log.Printf("From cache with id: %s", uid)
		return order, nil
	}
	order, err := postgres.GetOrderByUid(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	log.Printf("From db with id: %s", uid)

	err = InitCache()
	if err != nil {
		return nil, err
	}

	return order, nil
}
