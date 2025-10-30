package services

import (
	"context"
	"time"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
	"young-eagles/internal/dbmodels"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
)

type FlightService interface {
	PostFlight(ctx context.Context, pilot models.Pilot, child models.Child, plane models.Plane) (*models.Flight, error)
	PatchFlightFinished(ctx context.Context, flightUUID uuid.UUID) error
}

type flightServiceImpl struct {
	flightDao dao.FlightDao
	pilotDao  dao.PilotDao
	childDao  dao.ChildrenDao
}

func NewFlightService(fDao dao.FlightDao, pDao dao.PilotDao, cDao dao.ChildrenDao) FlightService {
	return &flightServiceImpl{
		flightDao: fDao,
		pilotDao:  pDao,
		childDao:  cDao,
	}
}

func (f flightServiceImpl) PostFlight(ctx context.Context, pilot models.Pilot, child models.Child, plane models.Plane) (*models.Flight, error) {
	/*
		Create db models for pilot and child
		Create flight information
		Post flight information to database
	*/
	fDBModel := dbmodels.FlightInformation{
		CreatedTS:       null.TimeFrom(time.Now()),
		Status:          null.StringFrom(models.InProgress.String()),
		PilotUUID:       pilot.PilotUuid.String(),
		ChildUUID:       null.StringFrom(child.UUID.String()),
		PlaneCallNumber: null.StringFrom(plane.CallNumber),
	}

	flightDB, err := f.flightDao.PutFlightDatum(ctx, fDBModel)
	if err != nil {
		return nil, err
	}

	flight := models.Flight{
		UUID:      uuid.MustParse(flightDB.UUID),
		PilotUUID: uuid.MustParse(flightDB.PilotUUID),
		ChildUUID: uuid.MustParse(flightDB.ChildUUID.String),
		Status:    models.FlightStatus(flightDB.Status.String),
	}
	return &flight, nil
}

func (f flightServiceImpl) PatchFlightFinished(ctx context.Context, flightUUID uuid.UUID) error {

	return nil
}
