package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"tap/cfg"
	"tap/internal/handlers"
	repo "tap/internal/repositories"
	"tap/internal/routes"
	services "tap/internal/services"
)

func main() {
	defer fmt.Println("Bye!")
	//Config
	err := cfg.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()
	app.Use(logger.New(logger.Config{
    Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
	//sss

	// Connect to MongoDB
	connectionUrl := cfg.DB()
	clientOptions := options.Client().ApplyURI(connectionUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(context.TODO())
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
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
	port := cfg.Port()
	log.Fatal(app.Listen(port))
}