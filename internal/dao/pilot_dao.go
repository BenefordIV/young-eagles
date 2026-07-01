package dao

import (
	"context"
	"time"
	"young-eagles/external/models"
	dbmodels "young-eagles/internal/db/gen/models"

	"github.com/aarondl/opt/omit"
	"github.com/gofrs/uuid/v5"
	"github.com/stephenafamo/bob"
)

type PilotDao interface {
	AddPilot(ctx context.Context, pilot models.Pilot) (*models.Pilot, error)
	GetPilotByNameChapterCombo(ctx context.Context, firstName, lastName string, eaaChapter int) (*models.Pilot, error)
	GetPilotByUUID(ctx context.Context, uuid uuid.UUID) (*models.Pilot, error)
	UpdatePilot(ctx context.Context, update *models.Pilot) error
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
	s := &dbmodels.PilotSetter{
		FirstName:  omit.From(pilot.PilotFirstName),
		LastName:   omit.From(pilot.PilotLastName),
		Email:      omit.From(pilot.PilotEmail),
		EaaChapter: omit.From(int32(pilot.EaaChapterNumber)),
		CreatedAt:  omit.From(time.Now()),
		UpdatedAt:  omit.From(time.Now()),
	}

	pi, err := dbmodels.Pilots.Insert(s).One(ctx, p.dbConn)
	if err != nil {
		return nil, err
	}

	added := models.PilotFromDb(*pi)

	return added, nil
}

func (p pilotDaoImpl) GetPilotByNameChapterCombo(ctx context.Context, firstName, lastName string, eaaChapter int) (*models.Pilot, error) {
	pilot, err := dbmodels.Pilots.Query(dbmodels.SelectWhere.Pilots.FirstName.EQ(firstName),
		dbmodels.SelectWhere.Pilots.LastName.EQ(lastName),
		dbmodels.SelectWhere.Pilots.EaaChapter.EQ(int32(eaaChapter))).
		One(ctx, p.dbConn)
	if err != nil {
		return nil, err
	}

	found := models.PilotFromDb(*pilot)

	return found, nil
}

func (p pilotDaoImpl) GetPilotByUUID(ctx context.Context, uuid uuid.UUID) (*models.Pilot, error) {
	pilot, err := dbmodels.Pilots.Query(dbmodels.SelectWhere.Pilots.ID.EQ(uuid)).One(ctx, p.dbConn)
	if err != nil {
		return nil, err
	}

	found := models.PilotFromDb(*pilot)

	return found, nil
}

func (p pilotDaoImpl) UpdatePilot(ctx context.Context, update *models.Pilot) error {
	existing, err := dbmodels.Pilots.Query(dbmodels.SelectWhere.Pilots.ID.EQ(update.PilotUuid)).One(ctx, p.dbConn)
	if err != nil {
		return err
	}

	s := &dbmodels.PilotSetter{
		FirstName:  omit.From(update.PilotFirstName),
		LastName:   omit.From(update.PilotLastName),
		Email:      omit.From(update.PilotEmail),
		EaaChapter: omit.From(int32(update.EaaChapterNumber)),
	}

	if err = existing.Update(ctx, p.dbConn, s); err != nil {
		return err
	}

	return nil
}
