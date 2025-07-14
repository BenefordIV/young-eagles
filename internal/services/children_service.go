package services

import (
	"context"
	"github.com/google/uuid"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
)

type ChildrenService interface {
	PostChildDatum(ctx context.Context, child models.Child) (*models.Child, error)
	GetChildren(ctx context.Context, parentUuid uuid.UUID) ([]models.Child, error)
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

func (c childrenServiceImpl) GetChildren(ctx context.Context, parentUuid uuid.UUID) ([]models.Child, error) {
	children, err := c.childDao.GetChildrenByParentID(ctx, parentUuid)
	if err != nil {
		return nil, err
	}

	cM := make([]models.Child, len(children))
	for i, child := range children {
		cM[i] = *models.ChildFromDb(*child)
	}

	return cM, nil
}
