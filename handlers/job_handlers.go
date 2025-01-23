package handlers

import (
	"net/http"
	"time"

	"github.com/adi790uu/kirana-club-assignment/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SubmitJob(db *gorm.DB) fiber.Handler {
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
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid request"})
		}
		if request.Count != len(request.Visits) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "count does not match visits length"})
		}

		job := models.Job{
			Status:    false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		db.Create(&job)

		for _, visit := range request.Visits {
			for _, imageURL := range visit.ImageURL {
				perimeter := ProcessImage(imageURL)
				image := models.Image{
					JobID:     job.ID,
					StoreID:   visit.StoreID,
					URL:       imageURL,
					Perimeter: perimeter,
					Status:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				db.Create(&image)
			}
		}

		job.Status = true
		job.UpdatedAt = time.Now()
		db.Save(&job)

		return c.Status(http.StatusCreated).JSON(fiber.Map{"success": true, "job_id": job.ID})
	}
}

func GetJobStatus(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jobID := c.Params("jobid")
		var job models.Job


		if err := db.First(&job, jobID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"success": false, "message": "job not found"})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success": true,
			"job_id":  job.ID,
			"status":  job.Status,
		})
	}
}
