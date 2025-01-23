package config

import (
	"log"
	"time"

	"github.com/adi790uu/kirana-club-assignment/models"
	"gorm.io/gorm"
)

func StartWorker(db *gorm.DB, jobQueue chan uint) {
	for jobID := range jobQueue {
		processJob(db, jobID)
	}
}

func processJob(db *gorm.DB, jobID uint) {
	log.Printf("Processing job ID: %d", jobID)

	var job models.Job
	if err := db.First(&job, jobID).Error; err != nil {
		log.Printf("Job ID %d not found: %v", jobID, err)
		return
	}

	var images []models.Image
	if err := db.Where("job_id = ?", jobID).Find(&images).Error; err != nil {
		log.Printf("Failed to fetch images for job ID %d: %v", jobID, err)
		return
	}

	for _, img := range images {
		time.Sleep(time.Duration(2 * time.Second))
		img.Perimeter = 2 * (100 + 50)                                  
		img.Status = true
		img.UpdatedAt = time.Now()
		db.Save(&img)
	}

	job.Status = true
	job.UpdatedAt = time.Now()
	db.Save(&job)

	log.Printf("Completed job ID: %d", jobID)
}
