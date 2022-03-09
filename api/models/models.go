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
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type GenericMessage struct {
	Message string `json:"message"`
}
