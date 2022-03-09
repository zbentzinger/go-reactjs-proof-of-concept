package main

import (
	"api/api/controllers"
	"api/api/database"
	"api/api/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func init() {
	database.Connect("sqlite")

	database.Connection.AutoMigrate(&models.Movie{})
	database.Connection.AutoMigrate(&models.User{})

	log.Println("Successfully Migrated database.")
}

func main() {

	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)

	basePathRouter := router.PathPrefix("/api/v1").Subrouter()
	basePathRouter.Use(loggingMiddleware)
	basePathRouter.HandleFunc("/healthcheck/", controllers.Healthcheck).Methods("GET")

	authRouter := basePathRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup/", controllers.Signup).Methods("POST")
	authRouter.HandleFunc("/login/", controllers.Login).Methods("POST")

	moviesRouter := basePathRouter.PathPrefix("/movies").Subrouter()
	moviesRouter.Use(authMiddleware)
	moviesRouter.HandleFunc("/", controllers.ListMovies).Methods("GET")
	moviesRouter.HandleFunc("/", controllers.CreateMovie).Methods("POST")
	moviesRouter.HandleFunc("/random/", controllers.GetRandomMovie).Methods("GET")
	moviesRouter.HandleFunc("/{id}/", controllers.GetMovie).Methods("GET")
	moviesRouter.HandleFunc("/{id}/", controllers.UpdateMovie).Methods("PUT")
	moviesRouter.HandleFunc("/{id}/", controllers.DeleteMovie).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)

		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if len(tokenString) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(
				models.GenericMessage{
					Message: "Missing Authorization Header",
				},
			)

			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		_, err := controllers.VerifyToken(tokenString)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(
				models.GenericMessage{
					Message: "Error verifying JWT token: " + err.Error(),
				},
			)

			return
		}

		next.ServeHTTP(w, r)

	})
}
