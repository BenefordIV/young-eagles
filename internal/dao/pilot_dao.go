package dao

import (
	"context"
	"young-eagles/external/models"

	"github.com/stephenafamo/bob"
)

type PilotDao interface {
	AddPilot(ctx context.Context, pilot models.Pilot) (*models.Pilot, error)
	GetPilotByNameChapterCombo(ctx context.Context, firstName, lastName string, eaaChapter int) (*models.Pilot, error)
	GetPilotByUUID(ctx context.Context, uuid string) (*models.Pilot, error)
	UpdatePilot(ctx context.Context, update *models.Pilot) (*models.Pilot, error)
}

type pilotDaoImpl struct {
	dbConn bob.DB
}

func NewPilotDao(conn bob.DB) PilotDao {
	return &pilotDaoImpl{
		dbConn: conn,
	}
}

func (p pilotDaoImpl) AddPilot(ctx context.Context, pilot models.Pilot) (*models.Pilot, error) {
	//log.Println("adding pilot to database")
	//log.Println(pilot.UUID)
	//err := pilot.Insert(ctx, p.dbConn.DbConn, boil.Blacklist(dbmodels.PilotDatumColumns.DeletedTS,
	//	dbmodels.PilotDatumColumns.UpdatedAt))
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &pilot, nil
	panic("implement me")
}

func (p pilotDaoImpl) GetPilotByNameChapterCombo(ctx context.Context, firstName, lastName string, eaaChapter int) (*models.Pilot, error) {
	//pilot, err := dbmodels.PilotData(dbmodels.PilotDatumWhere.PilotLastName.EQ(lastName),
	//	dbmodels.PilotDatumWhere.PilotFirstName.EQ(firstName),
	//	dbmodels.PilotDatumWhere.EaaChapterNumber.EQ(eaaChapter)).One(ctx, p.dbConn.DbConn)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return pilot, nil
	panic("implement me")
}

func (p pilotDaoImpl) GetPilotByUUID(ctx context.Context, uuid string) (*models.Pilot, error) {
	//pilot, err := dbmodels.PilotData(dbmodels.PilotDatumWhere.UUID.EQ(uuid)).One(ctx, p.dbConn.DbConn)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return pilot, nil
	panic("implement me")
}

func (p pilotDaoImpl) UpdatePilot(ctx context.Context, update *models.Pilot) (*models.Pilot, error) {
	//_, err := update.Update(ctx, p.dbConn.DbConn, boil.Infer())
	//if err != nil {
	//	return nil, err
	//}
	//
	//return update, nil
	panic("implement me")
}
