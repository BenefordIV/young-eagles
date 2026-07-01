package services

import (
	"context"
	"young-eagles/external/models"
	"young-eagles/internal/dao"

	"github.com/friendsofgo/errors"
	"github.com/gofrs/uuid/v5"
)

type ChildrenService interface {
	PostChildDatum(ctx context.Context, child models.Child) (*models.Child, error)
	GetChildByUUID(ctx context.Context, uuid uuid.UUID) (*models.Child, error)
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

	childDb, err := c.childDao.AddChildData(ctx, child)
	if err != nil {
		return nil, err
	}

	return childDb, nil
}

func (c childrenServiceImpl) GetChildByUUID(ctx context.Context, uuid uuid.UUID) (*models.Child, error) {
	ec, err := c.childDao.GetChildByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return ec, nil
}
