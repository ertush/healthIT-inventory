package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ertush/healthIT-inventory/App/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the Equipment by Id

	var equipment models.Equipment

	if result := h.DB.First(&equipment, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that Equipment
	h.DB.Delete(&equipment)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&map[string]string{
		"message" : "Deleted",
	})
}
