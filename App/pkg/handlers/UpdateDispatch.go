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

func (h handler) UpdateDispatch(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedDispatch models.Dispatch
	json.Unmarshal(body, &updatedDispatch)

	var dispatch models.Dispatch

	if result := h.DB.First(&dispatch, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	dispatch.Name = updatedDispatch.Name
	dispatch.Equipment = updatedDispatch.Equipment
	dispatch.Facility = updatedDispatch.Facility
	dispatch.DateIssued = updatedDispatch.DateIssued

	h.DB.Save(&dispatch)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&map[string]string{
		"message": "updated item",
	})
}
