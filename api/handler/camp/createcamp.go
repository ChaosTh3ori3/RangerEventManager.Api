package camp

import (
	"fmt"
	"net/http"
	"time"

	entities_camp "github.com/ChaosTh3ori3/RangerEventManager.Api/models/entities/camp"
	"github.com/ChaosTh3ori3/RangerEventManager.Api/persistens/repositories"
)

type CreateCampHandler struct {
	campRepo *repositories.CampRepository
}

func NewCreateCampHandler(campRepo *repositories.CampRepository) CreateCampHandler {
	return CreateCampHandler{
		campRepo: campRepo,
	}
}

// Handler1 ist die Funktion, die auf den Endpunkt "/api/endpoint1" reagiert.
func (cch CreateCampHandler) HandleCreateCamp(w http.ResponseWriter, r *http.Request) {
	// Hier kannst du die gewünschte Logik für deinen Endpunkt implementieren
	fmt.Fprintf(w, "Hallo von Handler1!")

	camp := entities_camp.Camp{
		Number:    1,
		Name:      "Test1234",
		StartDate: time.Now(),
		EndDate:   time.Now(),
	}

	cch.campRepo.CreateNewCamp(camp)
}
