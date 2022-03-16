package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h handler) Logout(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	w.Header().Add("Set-Cookie", cookie.String())
	json.NewEncoder(w).Encode(&map[string]string{
		"message": "success",
	})
}
