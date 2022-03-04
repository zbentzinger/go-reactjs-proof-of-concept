package controllers

import (
	"api/api/database"
	"api/api/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ListMovies(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)

	var movies []models.Movie

	database.Connection.Find(&movies)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(movies)

}

func GetMovie(w http.ResponseWriter, r *http.Request) {

	log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)

	params := mux.Vars(r)
	id := params["id"]
	var movie models.Movie

	result := database.Connection.First(&movie, id)

	if result.Error == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(movie)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)

	requestBody, _ := ioutil.ReadAll(r.Body)
	var movie models.Movie

	json.Unmarshal(requestBody, &movie)
	database.Connection.Create(&movie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)

	var movie models.Movie

	requestBody, _ := ioutil.ReadAll(r.Body)
	requestParams := mux.Vars(r)

	database.Connection.First(&movie, requestParams["id"])

	json.Unmarshal(requestBody, &movie)
	database.Connection.Save(&movie)

	w.WriteHeader(http.StatusNoContent)

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)

	var movie models.Movie

	requestBody, _ := ioutil.ReadAll(r.Body)
	requestParams := mux.Vars(r)

	json.Unmarshal(requestBody, &movie)
	database.Connection.Delete(&movie, requestParams["id"])

	w.WriteHeader(http.StatusNoContent)
}
