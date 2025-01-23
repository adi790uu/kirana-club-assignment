package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adi790uu/kirana-club-assignment/config"
	"github.com/adi790uu/kirana-club-assignment/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	app := fiber.New()

	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	log.Println("Connected to the database..")

	jobQueue := make(chan uint, 100)
	go config.StartWorker(db, jobQueue)

	routes.SetupRoutes(app, db, jobQueue)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Printf("Server running on port %s", port)

	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
