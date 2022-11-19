package controllers

import (
	"BudgBackend/src/config"
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

//var jwtToken = []byte("secret")
var jwtToken []byte

//obtengo la clave para generar el token
func init() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	jwtToken = []byte(cfg.JwtKey)

}

// CreateToken crea un token JWT
func CreateToken(w http.ResponseWriter, r *http.Request) {

	var user models.User
	// Get the JSON body and decode into credentials
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" || user.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Username and password are required"})
		return
	}

	okLogin, errLogin := user.ValidateLogin()
	if errLogin != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(responses.Exception{Message: errLogin.Error()})
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if okLogin {
		expiration := time.Now().Add(time.Hour * time.Duration(1)).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"id":       user.ID,
			"exp":      expiration,
		})
		tokenString, error := token.SignedString(jwtToken)
		if error != nil {
			fmt.Println(error)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		////create response with token and expiration
		//response := responses.TokenResponse{Token: tokenString, Expiration: expiration}
		//err := json.NewEncoder(w).Encode(response)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}

		json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString, Expiration: expiration})

	}
}
