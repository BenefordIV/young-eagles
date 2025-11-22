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
	CompleteFlight(ctx context.Context, flightUUID uuid.UUID) error
	GetFlight(ctx context.Context, flightUUID uuid.UUID) (*models.Flight, error)
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

	flight, err := f.flightDao.UpdateFlightDatum(ctx, fDBModel)
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (f flightServiceImpl) CompleteFlight(ctx context.Context, flightUUID uuid.UUID) error {
	fDbModel := dbmodels.FlightInformation{
		Status: null.StringFrom(models.Ended.String()),
		UUID:   flightUUID.String(),
	}

	_, err := f.flightDao.PutFlightDatum(ctx, fDbModel)

	if err != nil {
		return err
	}
	return nil
}

func (f flightServiceImpl) GetFlight(ctx context.Context, flightUUID uuid.UUID) (*models.Flight, error) {
	flightDb, err := f.flightDao.GetFlightData(ctx, flightUUID)

	if err != nil {
		return nil, err
	}

	flight := models.FlightFromDB(*flightDb)

	return &flight, nil
}
