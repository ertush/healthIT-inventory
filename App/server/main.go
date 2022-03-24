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

	// Books handlers
	router.HandleFunc("/api/books", h.AuthMiddleware(h.GetAllBooks)).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", h.AuthMiddleware(h.GetBook)).Methods(http.MethodGet)
	router.HandleFunc("/api/books", h.AuthMiddleware(h.AddBook)).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id}", h.AuthMiddleware(h.UpdateBook)).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{id}", h.AuthMiddleware(h.DeleteBook)).Methods(http.MethodDelete)

	//Equipment handlers
	router.HandleFunc("/api/equipment", h.AuthMiddleware(h.AddEquipment)).Methods(http.MethodPost)
	router.HandleFunc("/api/equipment/{id}", h.AuthMiddleware(h.GetEquipment)).Methods(http.MethodGet)
	router.HandleFunc("/api/equipment", h.AuthMiddleware(h.GetAllEquipment)).Methods(http.MethodGet)
	router.HandleFunc("/api/equipment/{id}", h.AuthMiddleware(h.DeleteEquipment)).Methods(http.MethodDelete)

	//Category handlers 

	// auth endpoints
	router.HandleFunc("/api/register", h.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/logout", h.Logout).Methods(http.MethodGet)
	router.HandleFunc("/api/user", h.User).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
