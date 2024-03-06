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

func GetAll(ctx context.Context) (orders []models.Order, err error) {
	result := DB.Find(&orders)
	if result.Error != nil {
		log.Fatalf("failed to find all orders: %v", result.Error.Error())
		return nil, result.Error
	}

	return orders, nil
}

func GetOrderByUid(ctx context.Context, uid string) (order models.Order, err error) {
	result := DB.Where("order_uid = ?", uid).Find(&order)
	if result.Error != nil {
		log.Fatalf("failed to find order by order_uid: %s, cause of: %v", uid, result.Error.Error())
	}
	return order, result.Error
}

func CreateOrder(ctx context.Context, order models.Order) (err error) {
	result := DB.Create(&order)
	if result.Error != nil {
		log.Printf("failed to create order: %v", result.Error.Error())
		return result.Error
	}
	log.Println("added to DB")
	return nil
}
