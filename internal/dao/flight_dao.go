package dao

import (
	"context"
	"log"
	"young-eagles/internal/dbmodels"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type FlightDao interface {
	PutFlightDatum(ctx context.Context, information dbmodels.FlightInformation) (*dbmodels.FlightInformation, error)
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
