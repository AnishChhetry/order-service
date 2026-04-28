package services

import (
	"context"
	"errors"
	"log"
	"order-service/models"
	"sync"
	"time"
)

type InMemService struct {
	mu     sync.RWMutex
	orders []models.Order
}

func NewInMemService(order []models.Order) *InMemService {
	return &InMemService{
		orders: order,
	}
}

func (s *InMemService) GetOrderList(ctx context.Context) ([]models.Order, error) {
	select {
	case <-time.After(100 * time.Millisecond):
		s.mu.RLock()
		defer s.mu.RUnlock()
		return s.orders, nil

	case <-ctx.Done():
		log.Println("GetOrderList canceled due to timeout")
		return nil, ctx.Err()
	}

}

func (s *InMemService) Create(ctx context.Context, order models.Order) (models.Order, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.orders {
		if s.orders[i].ID == order.ID {
			if s.orders[i].Item == order.Item {
				s.orders[i].Quantity += order.Quantity
				return s.orders[i], nil
			}
			log.Println("Error: Item with same ID already exists")
			return models.Order{}, errors.New("Item with same ID already exists")
		}

	}
	if order.Item == "" {
		log.Println("Error: Item Name is Required")
		return models.Order{}, errors.New("Item Name is Required")
	}
	s.orders = append(s.orders, order)

	return order, nil
}
