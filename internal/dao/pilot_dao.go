package dao

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"young-eagles/internal/dbmodels"
)

type PilotDao interface {
	AddPilot(ctx context.Context, pilot dbmodels.PilotDatum) (*dbmodels.PilotDatum, error)
	GetPilotByNameChapterCombo(ctx context.Context, firstName, lastName string, eaaChapter int) (*dbmodels.PilotDatum, error)
	GetPilotByUUID(ctx context.Context, uuid string) (*dbmodels.PilotDatum, error)
	UpdatePilot(ctx context.Context, update *dbmodels.PilotDatum) (*dbmodels.PilotDatum, error)
}

type pilotDaoImpl struct {
	dbConn DbConnection
}

func NewPilotDao(conn DbConnection) PilotDao {
	return &pilotDaoImpl{
		dbConn: conn,
	}
}

func (p pilotDaoImpl) AddPilot(ctx context.Context, pilot dbmodels.PilotDatum) (*dbmodels.PilotDatum, error) {
	log.Println("adding pilot to database")

	err := pilot.Insert(ctx, p.dbConn.DbConn, boil.Blacklist(dbmodels.PilotDatumColumns.DeletedTS,
		dbmodels.PilotDatumColumns.UpdatedAt))
	if err != nil {
		return nil, err
	}

	return &pilot, nil
}

func (p pilotDaoImpl) GetPilotByNameChapterCombo(ctx context.Context, firstName, lastName string, eaaChapter int) (*dbmodels.PilotDatum, error) {
	pilot, err := dbmodels.PilotData(dbmodels.PilotDatumWhere.PilotLastName.EQ(lastName),
		dbmodels.PilotDatumWhere.PilotFirstName.EQ(firstName),
		dbmodels.PilotDatumWhere.EaaChapterNumber.EQ(eaaChapter)).One(ctx, p.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return pilot, nil
}

func (p pilotDaoImpl) GetPilotByUUID(ctx context.Context, uuid string) (*dbmodels.PilotDatum, error) {
	pilot, err := dbmodels.PilotData(dbmodels.PilotDatumWhere.UUID.EQ(uuid)).One(ctx, p.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return pilot, nil
}

func (p pilotDaoImpl) UpdatePilot(ctx context.Context, update *dbmodels.PilotDatum) (*dbmodels.PilotDatum, error) {
	_, err := update.Update(ctx, p.dbConn.DbConn, boil.Infer())
	if err != nil {
		return nil, err
	}

	return update, nil
}
