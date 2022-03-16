package main

import (
	"log"
	"net/http"

	"github.com/ertush/healthIT-inventory/App/pkg/db"
	"github.com/ertush/healthIT-inventory/App/pkg/middleware"
	"github.com/ertush/healthIT-inventory/App/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/api/books", Middleware(http.HandleFunc(h.GetAllBooks)).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", Middleware(http.HandleFunc(h.GetBook)).Methods(http.MethodGet)
	router.HandleFunc("/api/books", Middleware(http.HandleFunc(h.AddBook)).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id}", Middleware(http.HandleFunc(h.UpdateBook)).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{id}", Middleware(http.HandleFunc(h.DeleteBook)).Methods(http.MethodDelete)

	// auth endpoints
	router.HandleFunc("/api/register", h.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/logout", h.Logout).Methods(http.MethodGet)
	router.HandleFunc("/api/user", h.User).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
