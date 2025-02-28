package services

import (
	"context"
	"errors"
	"github.com/volatiletech/null/v8"
	"time"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
	"young-eagles/internal/dbmodels"
)

type PilotService interface {
	PostPilotService(ctx context.Context, firstName, lastName, email string, eaaChapter int) (*models.Pilot, error)
}

type pilotServiceImpl struct {
	pilotDao dao.PilotDao
}

func NewPilotService(pilotDao dao.PilotDao) PilotService {
	return &pilotServiceImpl{
		pilotDao: pilotDao,
	}
}

func (p pilotServiceImpl) PostPilotService(ctx context.Context, firstName, lastName, email string, eaaChapter int) (*models.Pilot, error) {
	pilot, _ := p.pilotDao.GetPilotByNameChapterCombo(ctx, firstName, lastName, eaaChapter)

	if pilot != nil {
		return nil, errors.New("pilot already exists")
	}

	pDbModel := dbmodels.PilotDatum{
		PilotFirstName:   firstName,
		PilotLastName:    lastName,
		PilotEmail:       email,
		CreatedAt:        null.TimeFrom(time.Now()),
		EaaChapterNumber: eaaChapter,
		Status:           null.StringFrom(models.PilotStatusActive.ToString()),
	}

	dbModel, err := p.pilotDao.AddPilot(ctx, pDbModel)
	if err != nil {
		return nil, err
	}

	return models.PilotFromDb(*dbModel), nil
}
