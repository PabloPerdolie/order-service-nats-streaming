package postgres

import (
	"L0/order-service/internal/domain/models"
	"L0/order-service/pkg/client/postgres"
	"context"
	"fmt"
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
		log.Printf("failed to find order by order_uid: %v", result.Error.Error())
		return order, result.Error
	}

	if result.RowsAffected == 0 {
		log.Printf("order with order_uid: %s, doesnt exists", uid)
		return order, fmt.Errorf("record not found")
	}

	return order, nil
}

func CreateOrder(ctx context.Context, order models.Order) (err error) {
	result := DB.Create(&order)
	if result.Error != nil {
		return result.Error
	}
	log.Println("added to DB")
	return nil
}
