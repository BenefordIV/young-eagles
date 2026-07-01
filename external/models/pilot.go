package models

import (
	"young-eagles/internal/db/gen/models"

	"github.com/gofrs/uuid/v5"
)

type Pilot struct {
	PilotUuid        uuid.UUID `json:"pilotUuid"`
	PilotFirstName   string    `json:"pilotFirstName"`
	PilotLastName    string    `json:"pilotLastName"`
	PilotEmail       string    `json:"pilotEmail"`
	EaaChapterNumber int       `json:"eaaChapterNumber"`
}

func PilotFromDb(datum models.Pilot) *Pilot {
	return &Pilot{
		PilotUuid:        datum.ID,
		PilotFirstName:   datum.FirstName,
		PilotLastName:    datum.LastName,
		PilotEmail:       datum.Email,
		EaaChapterNumber: int(datum.EaaChapter),
	}
}

type PatchPilotBodyRequest struct {
	Pilot Pilot `json:"pilot"`
}

//func (r PatchPilotBodyRequest) GenerateUpdate(pilot *dbmodels.PilotDatum) (*dbmodels.PilotDatum, bool) {
//
//	changes := false
//
//	if r.Pilot.PilotFirstName != "" && r.Pilot.PilotFirstName != pilot.PilotFirstName {
//		pilot.PilotFirstName = r.Pilot.PilotFirstName
//		changes = true
//	}
//
//	if r.Pilot.PilotLastName != "" && r.Pilot.PilotLastName != pilot.PilotLastName {
//		pilot.PilotLastName = r.Pilot.PilotLastName
//		changes = true
//	}
//
//	if r.Pilot.PilotEmail != "" && r.Pilot.PilotEmail != pilot.PilotEmail {
//		pilot.PilotEmail = r.Pilot.PilotEmail
//		changes = true
//	}
//
//	if r.Pilot.EaaChapterNumber != 0 && r.Pilot.EaaChapterNumber != pilot.EaaChapterNumber {
//		pilot.EaaChapterNumber = r.Pilot.EaaChapterNumber
//		changes = true
//	}
//
//	return pilot, changes
//}
