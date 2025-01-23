package routes

import (
	"github.com/adi790uu/kirana-club-assignment/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/api/submit", handlers.SubmitJob(db))
	app.Get("/api/status", handlers.GetJobStatus(db))
}