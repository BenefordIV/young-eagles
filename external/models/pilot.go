package models

import (
	"github.com/google/uuid"
	"time"
	"young-eagles/internal/dbmodels"
)

type Pilot struct {
	PilotUuid        uuid.UUID
	PilotFirstName   string
	PilotLastName    string
	PilotEmail       string
	EaaChapterNumber int
	Status           PilotStatus
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type PilotStatus string

const (
	PilotStatusActive   PilotStatus = "ACTIVE"
	PilotStatusInactive PilotStatus = "INACTIVE"
	PilotStatusDeleted  PilotStatus = "DELETED"
)

func (p PilotStatus) ToString() string {
	return string(p)
}

func (p PilotStatus) Valid() bool {
	switch p {
	case PilotStatusActive, PilotStatusInactive, PilotStatusDeleted:
		return true
	default:
		return false
	}
}

func PilotFromDb(datum dbmodels.PilotDatum) *Pilot {
	status := PilotStatus(datum.Status.String)
	return &Pilot{
		PilotUuid:        uuid.MustParse(datum.UUID),
		PilotFirstName:   datum.PilotFirstName,
		PilotLastName:    datum.PilotLastName,
		PilotEmail:       datum.PilotEmail,
		EaaChapterNumber: datum.EaaChapterNumber,
		Status:           status,
		CreatedAt:        datum.CreatedAt.Time,
		UpdatedAt:        datum.UpdatedAt.Time,
		DeletedAt:        datum.DeletedTS.Time,
	}
}
