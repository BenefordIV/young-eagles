package models

import "github.com/google/uuid"

type Flight struct {
	UUID      uuid.UUID
	PilotUUID uuid.UUID
	ChildUUID uuid.UUID
	Status    FlightStatus
}

type FlightStatus string

const (
	InProgress FlightStatus = "in_progress"
	Ended      FlightStatus = "landed"
)

func (f FlightStatus) String() string {
	return string(f)
}

func (f FlightStatus) valid() bool {
	switch f {
	case InProgress, Ended:
		return true
	default:
		return false
	}
}
