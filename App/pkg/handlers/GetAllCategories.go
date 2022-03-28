package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ertush/healthIT-inventory/App/pkg/models"
)

func (h handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category

	if result := h.DB.Find(&categories); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
