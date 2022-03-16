package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/ertush/healthIT-inventory/App/pkg/models"
)

func (h handler) User(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("jwt")

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&map[string]string{
			"message": "unauthenticated",
		})
		return
	}

	token, err := jwt.ParseWithClaims(strings.Split(cookie.String(), "=")[1], &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(SecretKey), nil
	})

	// Unathenticated user

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&map[string]string{
			"message": "unauthenticated",
		})
		return
	}

	// Authenticated user

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	h.DB.Where("id = ?", claims.Issuer).First(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
