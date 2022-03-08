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
	router.Use(loggingMiddleware)

	router.HandleFunc("/v1/healthcheck/", controllers.Healthcheck).Methods("GET")

	router.HandleFunc("/v1/movies/", controllers.ListMovies).Methods("GET")
	router.HandleFunc("/v1/movies/", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/v1/movies/random/", controllers.GetRandomMovie).Methods("GET")
	router.HandleFunc("/v1/movies/{id}/", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/v1/movies/{id}/", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/v1/movies/{id}/", controllers.DeleteMovie).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)

		next.ServeHTTP(w, r)
	})
}
