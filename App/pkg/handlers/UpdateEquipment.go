package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/ertush/healthIT-inventory/App/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateEquipment(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedEquipment models.Equipment
	json.Unmarshal(body, &updatedEquipment)

	var equipment models.Equipment

	if result := h.DB.First(&equipment, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	equipment.Name = updatedEquipment.Name
	equipment.Category = updatedEquipment.Category
	equipment.Description = updatedEquipment.Description
	equipment.Category = updatedEquipment.Category
	equipment.Status = updatedEquipment.Status

	h.DB.Save(&equipment)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&map[string]string{
		"message": "updated item",
	})
}
