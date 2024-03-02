package postgres

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectToDB(ctx context.Context) (*gorm.DB, error) {
	dsn := os.Getenv("PSQL_AUTH")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	} else {
		log.Println("Successfully connected to database")
	}

	return db, nil
}
