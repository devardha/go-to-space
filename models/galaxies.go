package models

// Galaxy struct
type Galaxy struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"size:64"`
	Constellation string `json:"constellation" gorm:"size:128"`
	Type          string `json:"type" gorm:"size:64"`
}

// Galaxies struct
type Galaxies []Galaxy
