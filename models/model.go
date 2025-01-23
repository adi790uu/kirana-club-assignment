package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID        int       `db:"id" gorm:"primaryKey;autoIncrement"`
	Status    bool      `db:"status"`
	CreatedAt time.Time `db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `db:"updated_at" gorm:"autoUpdateTime"`

	Images []Image `gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Store struct {
	ID       string    `db:"id" gorm:"primaryKey"`
	Name     string `db:"name"`
	AreaCode string `db:"area_code"`
}

type Image struct {
	ID        int       `db:"id" gorm:"primaryKey;autoIncrement"`
	JobID     int       `db:"job_id" gorm:"not null"`
	StoreID   string       `db:"store_id"`
	Perimeter int       `db:"perimeter"`
	Status    bool    `db:"status"`
	URL       string    `db:"url"`
	CreatedAt time.Time `db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `db:"updated_at" gorm:"autoUpdateTime"`

	Job   Job   `gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Store Store `gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Job{}, &Store{}, &Image{})
	if err != nil {
		return err
	}
	return nil
}
