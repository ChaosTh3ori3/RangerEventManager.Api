package entities_camp

import (
	"time"

	"github.com/ChaosTh3ori3/RangerEventManager.Api/models/entities/general"
)

type Camp struct {
	Name               string
	Number             int
	StartDate          time.Time
	EndDate            time.Time
	Leaders            []general.Person
	Concept            string
	PreCampStartDate   time.Time
	PreCampEndDate     time.Time
	AfterCampStartDate time.Time
	AfterCampEndDate   time.Time
	Location           Location
	Members            []general.Person
	Tasks              []Task
	Sequences          []SequenceTask
}
