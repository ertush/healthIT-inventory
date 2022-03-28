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

	// Equipment handlers
	router.HandleFunc("/api/equipment", h.AuthMiddleware(h.AddEquipment)).Methods(http.MethodPost)
	router.HandleFunc("/api/equipment/{id}", h.AuthMiddleware(h.GetEquipment)).Methods(http.MethodGet)
	router.HandleFunc("/api/equipment", h.AuthMiddleware(h.GetAllEquipment)).Methods(http.MethodGet)
	router.HandleFunc("/api/equipment/{id}", h.AuthMiddleware(h.DeleteEquipment)).Methods(http.MethodDelete)

	// Category handlers
	router.HandleFunc("/api/category", h.AuthMiddleware(h.GetAllCategories)).Methods(http.MethodGet)
	router.HandleFunc("/api/category/{id}", h.AuthMiddleware(h.GetCategory)).Methods(http.MethodGet)
	router.HandleFunc("/api/category", h.AuthMiddleware(h.AddCategory)).Methods(http.MethodPost)
	router.HandleFunc("/api/category/{id}", h.AuthMiddleware(h.UpdateCategory)).Methods(http.MethodPut)
	router.HandleFunc("/api/category/{id}", h.AuthMiddleware(h.DeleteCategory)).Methods(http.MethodDelete)

	// Dispatch handlers
	router.HandleFunc("/api/dispatch", h.AuthMiddleware(h.GetAllDispatches)).Methods(http.MethodGet)
	router.HandleFunc("/api/dispatch/{id}", h.AuthMiddleware(h.GetDispatch)).Methods(http.MethodGet)
	router.HandleFunc("/api/dispatch", h.AuthMiddleware(h.AddDispatch)).Methods(http.MethodPost)
	router.HandleFunc("/api/dispatch/{id}", h.AuthMiddleware(h.UpdateDispatch)).Methods(http.MethodPut)
	router.HandleFunc("/api/dispatch/{id}", h.AuthMiddleware(h.DeleteDispatch)).Methods(http.MethodDelete)

	// Auth endpoints
	router.HandleFunc("/api/register", h.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/logout", h.Logout).Methods(http.MethodGet)
	router.HandleFunc("/api/user", h.User).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
