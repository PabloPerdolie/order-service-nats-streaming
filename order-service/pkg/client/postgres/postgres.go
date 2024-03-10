package postgres

import (
	"L0/order-service/internal/domain/config"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectToDB(ctx context.Context) (*gorm.DB, error) {
	conf := config.CONFIG.DB
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	} else {
		log.Println("Successfully connected to database")
	}

	return db, nil
}
