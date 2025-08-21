package models

import (
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"time"
	"young-eagles/internal/dbmodels"
)

type Child struct {
	UUID           uuid.UUID `json:"uuid"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	DateOfBirth    string    `json:"dateOfBirth"`
	HasCertificate bool      `json:"hasCertificate"`
}

func ChildFromDb(datum dbmodels.ChildInformation) *Child {
	return &Child{
		UUID:           uuid.MustParse(datum.UUID),
		FirstName:      datum.FirstName.String,
		LastName:       datum.LastName.String,
		DateOfBirth:    datum.DateOfBirth.Time.String(),
		HasCertificate: datum.HasCertificate.Bool,
	}
}

func (c *Child) ToDbModel() dbmodels.ChildInformation {
	dob, err := time.Parse(c.DateOfBirth, "2006-01-02")
	if err != nil {
		return dbmodels.ChildInformation{}
	}
	return dbmodels.ChildInformation{
		FirstName:      null.StringFrom(c.FirstName),
		LastName:       null.StringFrom(c.LastName),
		DateOfBirth:    null.TimeFrom(dob),
		HasCertificate: null.BoolFrom(c.HasCertificate),
	}
}
