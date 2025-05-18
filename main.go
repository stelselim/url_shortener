package main

import (
	"context"
	"log"
	"strconv"
	"url_shortener/controller"
	"url_shortener/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	port = 8080
)

func main() {
	ctx := context.Background()

	// Init Firebase Client
	_, err := service.GetFirestoreClient(ctx)
	if err != nil {
		log.Fatalf("Failed to get Firestore client: %v", err)
	}

	// Close Firebase Client
	defer func() {
		if err := service.CloseFirestoreClient(); err != nil {
			log.Printf("Error closing Firestore Client: %v", err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/shorten", controller.PostShortenController)
	e.GET("/:shortCode", controller.GetOriginalUrlController)
	e.GET("/stats/:shortCode", controller.GetShortenCodeStatsController)
	e.DELETE("/shorten/:shortCode", controller.DeleteShortenCodeController)

	var address string = ":" + strconv.Itoa(port)
	e.Logger.Fatal(e.Start(address))
}
