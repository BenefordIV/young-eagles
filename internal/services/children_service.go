package services

import (
	"context"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
)

type ChildrenService interface {
	PostChildDatum(ctx context.Context, child models.Child) (*models.Child, error)
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
	childDbModel := child.ToDbModel()

	childDb, err := c.childDao.PostChildData(ctx, childDbModel)
	if err != nil {
		return nil, err
	}

	return models.ChildFromDb(*childDb), nil
}
