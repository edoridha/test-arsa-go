package models

import "time"

type Film struct {
	ID                int             `json:"id" gorm:"primaryKey;autoIncrement"`
	Title             string          `json:"title" gorm:"type:varchar(255);not null"`
	Thumbnail         string          `json:"thumbnail" gorm:"type:varchar(255);not null"`
	ReleaseYear       string          `json:"release_year" gorm:"type:varchar(255);not null"`
	Director          string          `json:"director" gorm:"type:varchar(255);not null"`
	Sinopsis          string          `json:"sinopsis" gorm:"type:text;not null"`
	ProductionHouseID int             `json:"production_house_id" gorm:"not null;constraint:OnDelete:CASCADE;"`
	ProductionHouse   ProductionHouse `json:"production_house" gorm:"foreignKey:ProductionHouseID"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}
