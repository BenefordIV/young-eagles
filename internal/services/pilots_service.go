package services

import (
	"context"
	"database/sql"
	"errors"
	"young-eagles/external/models"
	"young-eagles/internal/dao"

	"github.com/gofrs/uuid/v5"
)

type PilotService interface {
	PostPilotData(ctx context.Context, pilot models.Pilot) (*models.Pilot, error)
	GetPilotData(ctx context.Context, pilotUuid uuid.UUID) (*models.Pilot, error)
	PatchUpdatePilotData(ctx context.Context, pilotUuid uuid.UUID, body models.PatchPilotBodyRequest) error
}

type pilotServiceImpl struct {
	pilotDao dao.PilotDao
}

func NewPilotService(pilotDao dao.PilotDao) PilotService {
	return &pilotServiceImpl{
		pilotDao: pilotDao,
	}
}

func (p pilotServiceImpl) PostPilotData(ctx context.Context, pilot models.Pilot) (*models.Pilot, error) {
	existingPilot, err := p.pilotDao.GetPilotByNameChapterCombo(ctx, pilot.PilotFirstName, pilot.PilotLastName, pilot.EaaChapterNumber)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if existingPilot != nil {
		return nil, errors.New("pilot already exists")
	}

	dbModel, err := p.pilotDao.AddPilot(ctx, pilot)
	if err != nil {
		return nil, err
	}

	return dbModel, nil
}

func (p pilotServiceImpl) GetPilotData(ctx context.Context, pilotUuid uuid.UUID) (*models.Pilot, error) {
	pilot, err := p.pilotDao.GetPilotByUUID(ctx, pilotUuid)

	if err != nil {
		return nil, err
	}

	return pilot, nil
}

func (p pilotServiceImpl) PatchUpdatePilotData(ctx context.Context, pilotUuid uuid.UUID, body models.PatchPilotBodyRequest) error {
	_, err := p.pilotDao.GetPilotByUUID(ctx, pilotUuid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	err = p.pilotDao.UpdatePilot(ctx, &body.Pilot)
	if err != nil {
		return err
	}
	return nil
}
