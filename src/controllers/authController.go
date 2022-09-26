package controllers

import (
	"BudgBackend/src/config"
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/http"
	"strings"
	"time"
)

//var jwtToken = []byte("secret")
var jwtToken []byte

func init() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	jwtToken = []byte(cfg.JwtKey)

}

// CreateToken creates a JWT token.
func CreateToken(w http.ResponseWriter, r *http.Request) {

	var user models.User
	// Get the JSON body and decode into credentials
	_ = json.NewDecoder(r.Body).Decode(&user)

	//user := models.User{}
	okLogin, errLogin := user.ValidateLogin()
	if errLogin != nil {
		err := json.NewEncoder(w).Encode(responses.Exception{Message: errLogin.Error()})
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if okLogin {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"id":       user.ID,
			"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		})
		tokenString, error := token.SignedString(jwtToken)
		if error != nil {
			fmt.Println(error)
		}
		json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString})

	}
}

// ValidateMiddleware validates the JWT token.
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
						w.WriteHeader(http.StatusUnauthorized)
					}
					return jwtToken, nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(responses.Exception{Message: error.Error()})
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid authorization token"})
					w.WriteHeader(http.StatusUnauthorized)
				}
			}
		} else {
			json.NewEncoder(w).Encode(responses.Exception{Message: "An authorization header is required"})
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

// ProtectedEndpoint is a protected endpoint.
func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	//obtener el token desde el header Authorization
	auth := r.Header.Get("Authorization")
	//separar el token del string "Bearer "
	bearerToken := strings.Split(auth, " ")[1]
	// validar el token
	token, _ := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtToken, nil
	})
	// si el token es valido, se obtiene el usuario del token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid authorization token"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}
}
