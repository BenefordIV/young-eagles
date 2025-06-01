package services

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"time"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
	"young-eagles/internal/dbmodels"
)

type ParentService interface {
	PostParentData(ctx context.Context, firstName, lastName, email, number string) (*models.Parent, error)
	GetParentData(ctx context.Context, parentUUID uuid.UUID) (*models.Parent, error)
}

type parentServiceImpl struct {
	parentDao dao.ParentDao
}

func NewParentService(parentDao dao.ParentDao) ParentService {
	return &parentServiceImpl{
		parentDao: parentDao,
	}
}

func (p parentServiceImpl) PostParentData(ctx context.Context, firstName, lastName, email, number string) (*models.Parent, error) {
	par, _ := p.parentDao.GetParentByEmail(ctx, email)

	if par != nil {
		return nil, errors.New("parent email exists in system, cannot add new parent with given email")
	}

	model := dbmodels.ParentInformation{
		UUID:        uuid.New().String(),
		FirstName:   null.StringFrom(firstName),
		LastName:    null.StringFrom(lastName),
		Email:       null.StringFrom(email),
		PhoneNumber: null.StringFrom(number),
		CreatedTS:   null.TimeFrom(time.Now()),
	}

	pDbM, err := p.parentDao.AddParentData(ctx, model)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add parent data")
	}

	pM := models.ParentFromDb(*pDbM)

	return pM, nil
}

func (p parentServiceImpl) GetParentData(ctx context.Context, parentUUID uuid.UUID) (*models.Parent, error) {
	par, err := p.parentDao.GetParentByID(ctx, parentUUID.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get parent data")
	}

	if par == nil {
		return nil, errors.New("parent not found")
	}

	pM := models.ParentFromDb(*par)

	return pM, nil
}
