package models

import (
	"github.com/google/uuid"
	"young-eagles/internal/dbmodels"
)

type Pilot struct {
	PilotUuid        uuid.UUID   `json:"pilotUuid"`
	PilotFirstName   string      `json:"pilotFirstName"`
	PilotLastName    string      `json:"pilotLastName"`
	PilotEmail       string      `json:"pilotEmail"`
	EaaChapterNumber int         `json:"eaaChapterNumber"`
	Status           PilotStatus `json:"status"`
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
	}
}

type PatchPilotBodyRequest struct {
	Pilot Pilot `json:"pilot"`
}

func (r PatchPilotBodyRequest) GenerateUpdate(pilot *dbmodels.PilotDatum) (*dbmodels.PilotDatum, bool) {

	changes := false

	if r.Pilot.PilotFirstName != "" && r.Pilot.PilotFirstName != pilot.PilotFirstName {
		pilot.PilotFirstName = r.Pilot.PilotFirstName
		changes = true
	}

	if r.Pilot.PilotLastName != "" && r.Pilot.PilotLastName != pilot.PilotLastName {
		pilot.PilotLastName = r.Pilot.PilotLastName
		changes = true
	}

	if r.Pilot.PilotEmail != "" && r.Pilot.PilotEmail != pilot.PilotEmail {
		pilot.PilotEmail = r.Pilot.PilotEmail
		changes = true
	}

	if r.Pilot.EaaChapterNumber != 0 && r.Pilot.EaaChapterNumber != pilot.EaaChapterNumber {
		pilot.EaaChapterNumber = r.Pilot.EaaChapterNumber
		changes = true
	}

	return pilot, changes
}
