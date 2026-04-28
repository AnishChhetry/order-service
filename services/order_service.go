package services

import (
	"context"
	"order-service/models"
)

type OrderService interface {
	GetOrderList(ctx context.Context) ([]models.Order, error)
	Create(ctx context.Context, order models.Order) (models.Order, error)
}
