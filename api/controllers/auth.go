package controllers

import (
	"api/api/database"
	"api/api/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func generateHash(s string) string {

	password := []byte(s)

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		log.Fatalln(err)
	}

	return string(hash)
}

func generateToken(user *models.User) string {

	expires := time.Minute * time.Duration(15)
	nowDateTime := time.Now().UTC()
	expiresDateTime := nowDateTime.Add(expires)

	signingKey := []byte("FooBar")

	claims := &jwt.StandardClaims{
		Audience:  "*.krs.home.arpa",
		ExpiresAt: expiresDateTime.Unix(),
		Id:        user.Email,
		IssuedAt:  nowDateTime.Unix(),
		Issuer:    "movieapi.krs.home.arpa",
		NotBefore: nowDateTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)

	if err != nil {
		log.Fatalln("Unable to generate JWT")
	}

	return signedToken
}

func VerifyToken(tokenString string) (jwt.Claims, error) {

	signingKey := []byte("FooBar")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err

}

func Signup(w http.ResponseWriter, r *http.Request) {

	var user models.User

	requestBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(requestBody, &user)

	user.Password = generateHash(user.Password)

	database.Connection.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(
		models.Token{
			Token: generateToken(&user),
		},
	)

}

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &user)

	plainPassword := user.Password

	var encryptedPassword string
	database.Connection.Model(&models.User{}).Where(&models.User{Username: user.Username}).Pluck("password", &encryptedPassword)

	encryptErr := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(plainPassword))

	if encryptErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(
			models.Token{
				Token: generateToken(&user),
			},
		)
	}

}
