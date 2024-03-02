package postgres

import (
	"L0/order-service/internal/domain/models"
	"L0/order-service/pkg/client/postgres"
	"context"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = postgres.ConnectToDB(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func GetOrderByUid(ctx context.Context, uid string) (order models.Order, err error) {

	return order, err
}

func CreateOrder(ctx context.Context, order models.Order) (err error) {
	//result := s.db.Create(&order.Delivery)
	//if result.Error != nil {
	//	log.Fatalf("failed to create delivery: %v", result.Error.Error())
	//}
	//result = s.db.Create(&order.Payment)
	//if result.Error != nil {
	//	log.Fatalf("failed to create delivery: %v", result.Error.Error())
	//}
	//for item := range order.Items {
	//	result = s.db.Create(&item)
	//	if result.Error != nil {
	//		log.Fatalf("failed to create delivery: %v", result.Error.Error())
	//	}
	//}
	//result = s.db.Create(&order.Items)
	//if result.Error != nil {
	//	log.Fatalf("failed to create delivery: %v", result.Error.Error())
	//}

	result := DB.Create(&order.Delivery)
	if result.Error != nil {
		log.Fatalf("failed to create order: %v", result.Error.Error())
	}
	return nil
}
