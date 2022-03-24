package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ertush/healthIT-inventory/App/pkg/models"
)

func (h handler) AddEquipment(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var equipment models.Equipment
	json.Unmarshal(body, &equipment)

	// Append to the equipments table
	if result := h.DB.Create(&equipment); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&map[string]string{
		"message": "Created",
	})
}
