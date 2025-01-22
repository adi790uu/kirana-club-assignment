package main

import (
	"log"
	"os"

	"github.com/adi790uu/kirana-club-assignment/database"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// func (r *Repository) SetupRoutes(app *fiber.App) {
// 	api := app.Group("/api")
// }

type Repository struct {
	DB    *gorm.DB
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := database.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load the database")
	}



	// repo := Repository{
	// 	DB:  db,
	// }

	app := fiber.New()
	// repo.SetupRoutes(app)

	log.Println("Server running on port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	
}
