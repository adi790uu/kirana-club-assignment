package handlers

import (
	"net/http"
	"time"

	"github.com/adi790uu/kirana-club-assignment/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SubmitJob(db *gorm.DB, jobQueue chan uint) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request struct {
			Count  int `json:"count"`
			Visits []struct {
				StoreID   string   `json:"store_id"`
				ImageURL  []string `json:"image_url"`
				VisitTime string   `json:"visit_time"`
			} `json:"visits"`
		}

		if err := c.BodyParser(&request); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
		}
		if request.Count != len(request.Visits) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "count does not match visits length"})
		}

		job := models.Job{
			Status:    "processing",
			StoreID:   request.Visits[0].StoreID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		db.Create(&job)

		for _, visit := range request.Visits {
			for _, imageURL := range visit.ImageURL {
				image := models.Image{
					JobID:     job.ID,
					StoreID:   visit.StoreID,
					URL:       imageURL,
					Status:    false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				db.Create(&image)
			}
		}

		jobQueue <- uint(job.ID)
		return c.Status(http.StatusCreated).JSON(fiber.Map{"job_id": job.ID})
	}
}


func GetJobStatus(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jobID := c.Query("jobid")
		if jobID == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "jobid is required"})
		}

		var job models.Job
		if err := db.Preload("Store").Preload("Images").First(&job, jobID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "job not found"})
		}

		if job.Status == "failed" {
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"job_id": job.ID,
				"status": job.Status,
				"error":  job.Error,
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"job_id": job.ID,
			"status": job.Status,
		})
	}
}
