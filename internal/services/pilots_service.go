package services

import (
	"context"
	"errors"
	"young-eagles/external/models"
	"young-eagles/internal/dao"

	"github.com/gofrs/uuid/v5"
)

type PilotService interface {
	PostPilotData(ctx context.Context, pilot models.Pilot) (*models.Pilot, error)
	GetPilotData(ctx context.Context, pilotUuid uuid.UUID) (*models.Pilot, error)
	PatchUpdatePilotData(ctx context.Context, pilotUuid uuid.UUID, body models.PatchPilotBodyRequest) (*models.Pilot, error)
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
	existingPilot, _ := p.pilotDao.GetPilotByNameChapterCombo(ctx, pilot.PilotFirstName, pilot.PilotLastName, pilot.EaaChapterNumber)

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

func (p pilotServiceImpl) PatchUpdatePilotData(ctx context.Context, pilotUuid uuid.UUID, body models.PatchPilotBodyRequest) (*models.Pilot, error) {
	//pilot, err := p.pilotDao.GetPilotByUUID(ctx, pilotUuid)
	//if err != nil {
	//	return nil, err
	//}
	//if pilot == nil {
	//	return nil, errors.New("pilot not found")
	//}
	//
	//pilotUpdate, updated := body.GenerateUpdate(pilot)
	//if updated {
	//	pilotDb, err := p.pilotDao.UpdatePilot(ctx, pilotUpdate)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return models.PilotFromDb(*pilotDb), nil
	//}
	//
	//return nil, errors.New("no changes made to pilot")
	panic("implement me")
}
