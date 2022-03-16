package middleware

import (
	"log"
	"net/http"
)

func middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")

		response, err := http.Get("/api/user")

		defer response.Body.Close()

		if err != nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}

		if response.StatusCode == http.StatusOK {
			next.ServeHTTP(w, r)
		}

	})
}
