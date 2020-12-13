package models

// Star - database model for star
type Star struct {
	ID            uint    `json:"id" gorm:"autoIncrement:true;primaryKey"`
	Name          string  `json:"name" gorm:"size:64"`
	Constellation string  `json:"constellation" gorm:"size:128"`
	Luminosity    float32 `json:"luminosity" gorm:"digits:32"`
	Temperature   float32 `json:"temperature" gorm:"digits:32"`
	Distance      float32 `json:"distance_ly" gorm:"digits:32"`
}

// Stars - database model for star list
type Stars []Star
