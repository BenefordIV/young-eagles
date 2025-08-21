package services

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"time"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
	"young-eagles/internal/dbmodels"
)

type ChildrenService interface {
	PostChildDatum(ctx context.Context, child models.Child) (*models.Child, error)
	GetChildByUUID(ctx context.Context, s string) (*models.Child, error)
}

type childrenServiceImpl struct {
	childDao dao.ChildrenDao
}

func NewChildrenService(childDao dao.ChildrenDao) ChildrenService {
	return &childrenServiceImpl{
		childDao: childDao,
	}
}

func (c childrenServiceImpl) PostChildDatum(ctx context.Context, child models.Child) (*models.Child, error) {
	ec, _ := c.childDao.GetChildByFirstLastName(ctx, child.FirstName, child.LastName)

	if ec != nil {
		return nil, errors.New("unable to post child as they already exist")
	}

	dob, err := time.Parse("2006-01-02", child.DateOfBirth)
	if err != nil {
		return nil, err
	}
	childDbModel := dbmodels.ChildInformation{
		FirstName:      null.StringFrom(child.FirstName),
		LastName:       null.StringFrom(child.LastName),
		DateOfBirth:    null.TimeFrom(dob),
		HasCertificate: null.BoolFrom(child.HasCertificate),
		CreatedTS:      null.TimeFrom(time.Now()),
	}

	childDb, err := c.childDao.PostChildData(ctx, childDbModel)
	if err != nil {
		return nil, err
	}

	return models.ChildFromDb(*childDb), nil
}

func (c childrenServiceImpl) GetChildByUUID(ctx context.Context, uuid string) (*models.Child, error) {
	ec, err := c.childDao.GetChildByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	ecM := models.ChildFromDb(*ec)

	return ecM, nil
}
