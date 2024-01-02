package entities_camp

import "time"

type SequenceTask struct {
	Task
	StartDate time.Time
	EndDate   time.Time
}
