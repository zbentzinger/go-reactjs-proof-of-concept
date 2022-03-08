package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	ReleaseDate string `json:"releaseDate"`
}

type Healthcheck struct {
	Status string `json:"status"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
