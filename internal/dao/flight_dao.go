package dao

import (
	"context"
	"log"
	"young-eagles/external/models"
	"young-eagles/internal/dbmodels"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type FlightDao interface {
	PutFlightDatum(ctx context.Context, information dbmodels.FlightInformation) (*dbmodels.FlightInformation, error)
	UpdateFlightDatum(ctx context.Context, information dbmodels.FlightInformation) (*models.Flight, error)
	GetFlightData(ctx context.Context, uuid uuid.UUID) (*dbmodels.FlightInformation, error)
}

type flightDao struct {
	dbConn DbConnection
}

func NewFlightDao(conn DbConnection) FlightDao {
	return &flightDao{
		dbConn: conn,
	}
}

func (f *flightDao) PutFlightDatum(ctx context.Context, information dbmodels.FlightInformation) (*dbmodels.FlightInformation, error) {
	log.Printf("posting flight data for pilot %s and child %s\n", information.PilotUUID, information.ChildUUID.String)
	err := information.Insert(ctx, f.dbConn.DbConn, boil.Blacklist(dbmodels.FlightInformationColumns.DeletedTS, dbmodels.FlightInformationColumns.UpdatedTS))

	if err != nil {
		return nil, err
	}

	return &information, nil
}

func (f *flightDao) UpdateFlightDatum(ctx context.Context, information dbmodels.FlightInformation) (*models.Flight, error) {
	log.Printf("ending flight for flight %s\n", information.UUID)

	_, err := information.Update(ctx, f.dbConn.DbConn, boil.Infer())
	if err != nil {
		return nil, err
	}

	fDb, err := f.GetFlightData(ctx, uuid.MustParse(information.UUID))

	if err != nil {
		return nil, err
	}

	flight := models.FlightFromDB(*fDb)

	return &flight, err
}
func (f *flightDao) GetFlightData(ctx context.Context, uuid uuid.UUID) (*dbmodels.FlightInformation, error) {
	flight, err := dbmodels.FlightInformations(dbmodels.FlightInformationWhere.UUID.EQ(uuid.String())).One(ctx, f.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return flight, nil
}
