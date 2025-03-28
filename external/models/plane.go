package models

import "young-eagles/internal/dbmodels"

type Plane struct {
	CallNumber string `json:"callNumber"`
	PlaneModel string `json:"planeModel"`
	PlaneMake  string `json:"planeMake"`
}

func PlaneFromDb(information dbmodels.PlaneInformation) Plane {
	return Plane{
		CallNumber: information.CallNumber,
		PlaneModel: information.Model.String,
		PlaneMake:  information.Model.String,
	}
}
