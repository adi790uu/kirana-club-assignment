package config

import (
	"encoding/json"
	"log"
	"time"

	"github.com/adi790uu/kirana-club-assignment/models"
	"github.com/adi790uu/kirana-club-assignment/utils"
	"gorm.io/gorm"
)

func StartWorker(db *gorm.DB, jobQueue chan uint) {
	for jobID := range jobQueue {
		processJob(db, jobID)
	}
}

func processJob(db *gorm.DB, jobID uint) {
	log.Printf("Starting processing for job ID: %d", jobID)

	var job models.Job
	if err := db.First(&job, jobID).Error; err != nil {
		log.Printf("Error: Job ID %d not found. %v", jobID, err)
		return
	}

	var images []models.Image
	if err := db.Where("job_id = ?", jobID).Find(&images).Error; err != nil {
		log.Printf("Error: Failed to fetch images for job ID %d. %v", jobID, err)
		return
	}

	var processingErrors []map[string]string

	for _, img := range images {
		log.Printf("Processing image ID: %d (URL: %s)", img.ID, img.URL)

		if !isValidStoreID(db, img.StoreID) {
			log.Printf("Error: Invalid store_id %s for image ID %d", img.StoreID, img.ID)
			processingErrors = append(processingErrors, map[string]string{
				"store_id": img.StoreID,
				"error":    "Invalid store_id",
			})
			continue
		}

		perimeter := utils.ProcessImage(img.URL)
		if perimeter == -1 {
			log.Printf("Error: Failed to process image ID: %d (URL: %s)", img.ID, img.URL)
			processingErrors = append(processingErrors, map[string]string{
				"store_id": img.StoreID,
				"error":    "Image processing failed",
			})

			img.Status = false
		} else {
			img.Status = true
			img.Perimeter = perimeter
		}

		img.UpdatedAt = time.Now()
		db.Save(&img)
	}

	job.UpdatedAt = time.Now()
	if len(processingErrors) > 0 {
		job.Status = "failed"
		jsonErrors, err := json.Marshal(processingErrors)
		if err != nil {
			log.Printf("Error: Failed to serialize errors: %v", err)
		}

		job.Error = jsonErrors
		log.Printf("Job ID %d completed with errors: %v", jobID, processingErrors)
	} else {
		job.Status = "completed"
		log.Printf("Job ID %d processed successfully", jobID)
	}

	db.Save(&job)
}


func isValidStoreID(db *gorm.DB, storeID string) bool {
	var store models.Store
	if err := db.Where("id = ?", storeID).First(&store).Error; err != nil {
		return false
	}
	return true
}