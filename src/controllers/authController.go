package controllers

import (
	"BudgBackend/src/config"
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
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
		ErrorLogger.Println(errLogin.Error())
		err := json.NewEncoder(w).Encode(responses.Exception{Message: "Error validando el usuario"})
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
			ErrorLogger.Println(error.Error())

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString, Expiration: expiration})

	}
}
