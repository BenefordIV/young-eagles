package models

import (
	"time"
	dbmodels "young-eagles/internal/db/gen/models"

	"github.com/gofrs/uuid/v5"
)

type Child struct {
	UUID           uuid.UUID `json:"uuid"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	HasCertificate bool      `json:"hasCertificate"`
}

func ChildFromDb(datum dbmodels.Child) *Child {
	return &Child{
		UUID:           datum.ID,
		FirstName:      datum.FirstName,
		LastName:       datum.LastName,
		DateOfBirth:    datum.DateOfBirth,
		HasCertificate: datum.HasCertificate,
	}
}

func (c *Child) ToDbModel() dbmodels.Child {
	return dbmodels.Child{
		FirstName:      c.FirstName,
		LastName:       c.LastName,
		DateOfBirth:    c.DateOfBirth,
		HasCertificate: c.HasCertificate,
	}
}
