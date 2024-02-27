package postgres

import (
	"L0/order-service/internal/domain/models"
	"context"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func init() {

}

func NewClient() {

}

func (s *Storage) GetOrderByUid(ctx context.Context, uid string) (order models.Order, err error) {
	return order, err
}

func (s *Storage) CreateOrder(ctx context.Context, order models.Order) (err error) {
	return nil
}
