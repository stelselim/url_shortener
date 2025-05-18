package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"url_shortener/controller"
	"url_shortener/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
)

const (
	port = 8080
)

func main() {
	ctx := context.Background()

	initFirebase(ctx)
	defer closeFirebase()

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/shorten", controller.PostShortenController)
	e.GET("/:shortCode", controller.GetOriginalUrlController)
	e.GET("/stats/:shortCode", controller.GetShortenCodeStatsController)
	e.DELETE("/shorten/:shortCode", controller.DeleteShortenCodeController)

	var address string = ":" + strconv.Itoa(port)

	// Start server and check error without os.Exit
	if err := e.Start(address); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

func initFirebase(ctx context.Context) {
	// Load Env variables
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	// Get Environment Variable for Firebase
	credPath := os.Getenv("FIREBASE_CREDENTIALS")
	if credPath == "" {
		log.Fatalf("Failed to get Firebase Credentials")
	}

	// Get Credentials File
	opt := option.WithCredentialsFile(credPath)
	// Init Firebase Client
	_, firestoreClientErr := service.GetFirestoreClient(ctx, opt)

	if firestoreClientErr != nil {
		log.Fatalf("Failed to get Firestore client: %v", firestoreClientErr)
	}
}

func closeFirebase() {
	if err := service.CloseFirestoreClient(); err != nil {
		log.Printf("Error closing Firestore Client: %v", err)
	}
}
