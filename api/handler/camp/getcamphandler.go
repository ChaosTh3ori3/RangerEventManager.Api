package camp

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChaosTh3ori3/RangerEventManager.Api/persistens/repositories"
	"github.com/gorilla/mux"
)

type GetCampHandler struct {
	campRepo *repositories.CampRepository
}

func NewGetCampHandler(campRepo *repositories.CampRepository) GetCampHandler {
	return GetCampHandler{
		campRepo: campRepo,
	}
}

func (gch GetCampHandler) GetCampByCampNumber(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	campnumberstring := vars["campNumber"]

	campnumber, err := strconv.Atoi(campnumberstring)
	if err != nil {
		http.Error(w, "Ungültiger Routenparameter 'id'", http.StatusBadRequest)
	}

	camp := gch.campRepo.GetCampByCampNumber(campnumber)

	jsonData, err := json.Marshal(camp)
	if err != nil {
		http.Error(w, "Fehler bei der JSON-Marshalling", http.StatusInternalServerError)
		return
	}

	// Setze den Content-Type-Header auf "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Schreibe die JSON-Daten zurück zum Client
	w.Write(jsonData)
}
