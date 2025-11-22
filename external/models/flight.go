package models

import (
	"young-eagles/internal/dbmodels"

	"github.com/google/uuid"
)

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

func FlightFromDB(datum dbmodels.FlightInformation) Flight {
	return Flight{
		UUID:      uuid.MustParse(datum.UUID),
		PilotUUID: uuid.MustParse(datum.PilotUUID),
		ChildUUID: uuid.MustParse(datum.ChildUUID.String),
		Status:    FlightStatus(datum.Status.String),
	}
}
