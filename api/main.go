package main

import (
	"api/api/controllers"
	"api/api/database"
	"api/api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func init() {
	database.Connect()

	database.Connection.AutoMigrate(&models.Movie{})
	log.Println("Successfully Migrated database.")
}

func main() {

	log.Println("Starting the HTTP server on port 8080")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/v1/movies/", controllers.ListMovies).Methods("GET")
	router.HandleFunc("/v1/movies/", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/v1/movies/{id}/", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/v1/movies/{id}/", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/v1/movies/{id}/", controllers.DeleteMovie).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8080", router))

}
