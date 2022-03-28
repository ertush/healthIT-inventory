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

func (h handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedCategory models.Category
	json.Unmarshal(body, &updatedCategory)

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	category.Name = updatedCategory.Name
	category.Description = updatedCategory.Description

	h.DB.Save(&category)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&map[string]string{
		"message": "updated item",
	})
}
