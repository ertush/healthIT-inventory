package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ertush/healthIT-inventory/App/pkg/models"
)

func (h handler) GetAllEquipment(w http.ResponseWriter, r *http.Request) {
	var equipments []models.Equipment

	if result := h.DB.Find(&equipments); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(equipments)
}
