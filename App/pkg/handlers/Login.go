package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/ertush/healthIT-inventory/App/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &data)

	var user models.User

	// Query first user
	h.DB.Where("email = ?", data["email"]).First(&user)

	// User not found
	if user.Id == 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			&map[string]string{
				"message": "user not found",
			})
		return
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&map[string]string{
			"message": "incorrect password",
		})
		return
	}

	extime := &jwt.Time{time.Now().Add(time.Hour * 24).UTC()} //1 day

	// Set jwt claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: extime,
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(&map[string]string{
			"message": "could not login",
		})
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	w.Header().Add("Set-Cookie", cookie.String())
	json.NewEncoder(w).Encode(&map[string]string{
		"message": "success",
	})

}
