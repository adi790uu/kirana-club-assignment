package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/adi790uu/kirana-club-assignment/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	
	dsn := "host=ep-aged-boat-a7sze0hr.ap-southeast-2.aws.neon.tech user=neondb_owner password=npg_d4bs7gRpLoUC dbname=neondb port=5432 sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}


	file, err := os.Open("/Users/adi790u/Desktop/kirana-club-assignment/StoreMasterAssignment.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()


	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		areaCode := record[0]
		storeName := record[1]
		storeID := record[2]

		store := models.Store{
			ID:       storeID, 
			Name:     storeName,
			AreaCode: areaCode,
		}

		if err := db.Create(&store).Error; err != nil {
			log.Printf("Failed to insert record for StoreID %s: %v", storeID, err)
		}
	}

	log.Println("Store data successfully inserted.")
}
