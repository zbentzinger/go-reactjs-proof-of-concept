package controllers

import (
	"api/api/database"
	"api/api/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var movie models.Movie
var movies []models.Movie

func ListMovies(w http.ResponseWriter, r *http.Request) {

	database.Connection.Find(&movies)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(movies)

}

func GetMovie(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	result := database.Connection.First(&movie, id)

	if result.Error == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(movie)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func GetRandomMovie(w http.ResponseWriter, r *http.Request) {

	result := database.Connection.Raw("SELECT * FROM movies ORDER BY RAND() LIMIT 1").Scan(&movie) // GORM has no Rand() function.

	if result.Error == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(movie)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	requestBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(requestBody, &movie)
	database.Connection.Create(&movie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

	requestBody, _ := ioutil.ReadAll(r.Body)
	requestParams := mux.Vars(r)

	database.Connection.First(&movie, requestParams["id"])

	json.Unmarshal(requestBody, &movie)
	database.Connection.Save(&movie)

	w.WriteHeader(http.StatusNoContent)

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

	requestBody, _ := ioutil.ReadAll(r.Body)
	requestParams := mux.Vars(r)

	json.Unmarshal(requestBody, &movie)
	database.Connection.Delete(&movie, requestParams["id"])

	w.WriteHeader(http.StatusNoContent)
}
