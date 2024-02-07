package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"tap/internal/handlers"
	repo "tap/internal/repositories"
	"tap/internal/routes"
	services "tap/internal/services"
)

func main() {
	// Create a Fiber app
	app := fiber.New()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://ethanham:a1a2b3b4@app.o2dc26k.mongodb.net/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Create a collection
	database := client.Database("test1")

	// Create a repository
	repository := repo.NewRepository(database)

	// Create a service
	service := services.NewService(repository)

	// Create a handler
	handler := handlers.NewHandler(service)

	// Define routes
	routes.SetupRoutes(app, handler)

	// Start server
	log.Fatal(app.Listen(":3000"))
}