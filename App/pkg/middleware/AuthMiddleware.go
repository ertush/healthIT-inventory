package middleware

import (
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")

		response, err := http.Get("/api/user")

		if err != nil || response.StatusCode != http.StatusOK {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}

		defer response.Body.Close()

		next.ServeHTTP(w, r)

	})
}
