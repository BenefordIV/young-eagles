package models

import (
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"young-eagles/internal/dbmodels"
)

type Parent struct {
	ID          uuid.UUID    `json:"id"`
	FirstName   *string      `json:"firstName"`
	LastName    *string      `json:"lastName"`
	Email       *string      `json:"email"`
	PhoneNumber string       `json:"phoneNumber"`
	Status      ParentStatus `json:"status"`
}

type ParentWithChildren struct {
	Par      Parent  `json:"parent"`
	Children []Child `json:"children"`
}

type ParentStatus string

const (
	Active   ParentStatus = "ACTIVE"
	Inactive ParentStatus = "INACTIVE"
	Deleted  ParentStatus = "DELETED"
)

func (p *ParentStatus) ToString() string {
	return string(*p)
}

func (p *ParentStatus) Valid() bool {
	switch *p {
	case Active, Inactive, Deleted:
		return true
	default:
		return false
	}
}

func ParentFromDb(datum dbmodels.ParentInformation) *Parent {
	status := ParentStatus(datum.Status.String)
	return &Parent{
		ID:          uuid.MustParse(datum.UUID),
		FirstName:   &datum.FirstName.String,
		LastName:    &datum.LastName.String,
		Email:       &datum.Email.String,
		PhoneNumber: datum.PhoneNumber.String,
		Status:      status,
	}
}

func (p *Parent) ToDbModel() dbmodels.ParentInformation {
	return dbmodels.ParentInformation{
		UUID:        p.ID.String(),
		FirstName:   null.StringFrom(*p.FirstName),
		LastName:    null.StringFrom(*p.LastName),
		Email:       null.StringFrom(*p.Email),
		PhoneNumber: null.StringFrom(p.PhoneNumber),
		Status:      null.StringFrom(p.Status.ToString()),
	}
}
