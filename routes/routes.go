package routes

import (
	"github.com/adi790uu/kirana-club-assignment/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jobQueue chan uint) {
	api := app.Group("/api")
	api.Post("/submit", handlers.SubmitJob(db, jobQueue))
	api.Get("/status", handlers.GetJobStatus(db))
}