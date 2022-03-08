package controllers

import (
	"api/api/models"
	"encoding/json"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {

	// Do some special logic to determine if the api is available.
	healthCheck := models.Healthcheck{Status: "healthy"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(healthCheck)
}
