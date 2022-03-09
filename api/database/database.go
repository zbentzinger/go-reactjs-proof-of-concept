package database

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var Connection *gorm.DB

var GetConnectionStringMysql = func(config Config) string {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.User, config.Password, config.ServerName, config.DB,
	)
	return connectionString
}

var GetConnectionStringSqlite = func() string {
	connectionString := "file::memory:?cache=shared"
	return connectionString
}

func connectMysql() {
	config := Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "toor",
		DB:         "movies",
	}

	db, err := gorm.Open(
		mysql.Open(
			GetConnectionStringMysql(config),
		),
		&gorm.Config{},
	)

	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connected to MySQL database")
		Connection = db
	}
}

func connectSqlite() {
	db, err := gorm.Open(
		sqlite.Open(
			GetConnectionStringSqlite(),
		),
		&gorm.Config{},
	)

	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connected to SQLite database")
		Connection = db
	}
}

func Connect(dbType string) {

	dbType = strings.ToLower(dbType)

	switch dbType {
	case "sqlite":
		connectSqlite()
	case "mysql":
		fallthrough
	default:
		connectMysql()
	}

}
