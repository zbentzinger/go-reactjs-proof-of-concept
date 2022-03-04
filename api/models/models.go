package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	ReleaseDate string `json:"releaseDate"`
}
