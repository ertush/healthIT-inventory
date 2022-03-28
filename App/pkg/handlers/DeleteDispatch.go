package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ertush/healthIT-inventory/App/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteDispatch(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the Dispatch by Id

	var dispatch models.Dispatch

	if result := h.DB.First(&dispatch, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that Dispatch
	h.DB.Delete(&dispatch)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&map[string]string{
		"message": "Deleted",
	})
}
