package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"tap/config"
	"tap/internal/handlers"
	repo "tap/internal/repositories"
	"tap/internal/routes"
	services "tap/internal/services"
)

func main() {
	//Config
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()
	app.Use(logger.New(logger.Config{
    Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
	app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:9000",
		AllowHeaders: "*",
		ExposeHeaders: "*",
		MaxAge: 24 * 60 * 60,
		AllowCredentials: true,
		AllowMethods: "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	app.Use (func(c *fiber.Ctx) error {
		c.Set("X-Powered-By", "Fiber")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	// Connect to MongoDB
	connectionUrl := config.ConnectionUrl()
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
	database := client.Database(config.DBName())

	// Create a repository
	repository := repo.NewRepository(database)

	// Create a service
	service := services.NewService(repository)

	// Create a handler
	handler := handlers.NewHandler(service)

	// Define routes
	routes.SetupRoutes(app, handler)

	// Start server
	port := config.Port()
	log.Fatal(app.Listen(port))
}