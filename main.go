package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"order-service/handlers"
	"order-service/models"
	"order-service/services"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var port string
	var serviceName string
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using default variables")
	}
	port = os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	serviceName = os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "order-service"
	}

	router := gin.Default()

	orderChan := make(chan models.Order, 5)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go services.FulfillmentWorker(ctx, orderChan)

	service := services.NewInMemService([]models.Order{
		{
			ID:       "101",
			Item:     "Laptop",
			Quantity: 1,
		},
		{
			ID:       "102",
			Item:     "Monitor",
			Quantity: 4,
		},
	})
	handler := &handlers.OrderHandler{
		Service:   service,
		OrderChan: orderChan,
	}

	router.GET("/health", handler.HealthStatus)
	router.GET("/orders", handler.OrderList)
	router.POST("/orders", handler.CreateOrder)

	log.Printf("Starting %s on port %s\n", serviceName, port)
	log.Printf("Health: http://localhost:%s/health\n", port)
	log.Printf("Orders: http://localhost:%s/orders", port)

	// log.Fatal(router.Run(":" + port))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		log.Printf("Starting %s on port %s\n", serviceName, port)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("\nReceived shutdown signal! Shutting down server...")

	cancel()

	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeout()

	err := srv.Shutdown(ctxTimeout)
	if err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server Exiting...")

}
