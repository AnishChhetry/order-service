package services

import (
	"context"
	"log"
	"order-service/models"
	"time"
)

func FulfillmentWorker(ctx context.Context, orderChan <-chan models.Order) {
	log.Println("Fulfillment worker started...")
	for {
		select {
		case order := <-orderChan:
			log.Printf("Worker: Processing order %s for item %s...\n", order.ID, order.Item)
			time.Sleep(1 * time.Second)
			log.Printf("Worker: Order %s fulfilled\n", order.ID)
		case <-ctx.Done():
			log.Println("Worker: Shutting down gracefully...")
			return
		}
	}
}
