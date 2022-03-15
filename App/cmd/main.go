package main

import (
	"log"
	"net/http"

	"github.com/ertush/healthIT-inventory/App/pkg/db"
	"github.com/ertush/healthIT-inventory/App/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/api/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/api/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{id}", h.DeleteBook).Methods(http.MethodDelete)
	router.HandleFunc("/api/register", h.Register).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
