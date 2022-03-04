package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var Connection *gorm.DB

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.User, config.Password, config.ServerName, config.DB,
	)
	return connectionString
}

func Connect() {

	config := Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "toor",
		DB:         "movies",
	}

	connectionString := GetConnectionString(config)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connected to MySQL database")
		Connection = db
	}

}
