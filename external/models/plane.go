package models

import "young-eagles/internal/db/gen/models"

type Plane struct {
	CallNumber string `json:"callNumber"`
	PlaneModel string `json:"planeModel"`
	PlaneMake  string `json:"planeMake"`
}

func PlaneFromDb(information models.Plane) *Plane {
	return &Plane{
		CallNumber: information.CallNumber,
		PlaneModel: information.Model,
		PlaneMake:  information.Make,
	}
}
