package services

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/google/uuid"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
)

type ParentService interface {
	PostParentData(ctx context.Context, parent models.Parent) (*models.Parent, error)
	GetParentData(ctx context.Context, parentUUID uuid.UUID) (*models.Parent, []models.Child, error)
	PostParentWithChildData(ctx context.Context, parent models.Parent, kid []models.Child) (*models.ParentWithChildren, error)
}

type parentServiceImpl struct {
	parentDao    dao.ParentDao
	childService ChildrenService
}

func NewParentService(parentDao dao.ParentDao, childrenService ChildrenService) ParentService {
	return &parentServiceImpl{
		parentDao:    parentDao,
		childService: childrenService,
	}
}

func (p parentServiceImpl) PostParentData(ctx context.Context, parent models.Parent) (*models.Parent, error) {
	par, _ := p.parentDao.GetParentByEmail(ctx, *parent.Email)

	if par != nil {
		return nil, errors.New("parent email exists in system, cannot add new parent with given email")
	}

	model := parent.ToDbModel()

	pDbM, err := p.parentDao.AddParentData(ctx, model)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add parent data")
	}

	pM := models.ParentFromDb(*pDbM)

	return pM, nil
}

func (p parentServiceImpl) GetParentData(ctx context.Context, parentUUID uuid.UUID) (*models.Parent, []models.Child, error) {
	par, err := p.parentDao.GetParentByID(ctx, parentUUID.String())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get parent data")
	}

	if par == nil {
		return nil, nil, errors.New("parent not found")
	}

	pM := models.ParentFromDb(*par)

	cM, err := p.childService.GetChildren(ctx, parentUUID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get children")
	}

	return pM, cM, nil
}

func (p parentServiceImpl) PostParentWithChildData(ctx context.Context, parent models.Parent, kid []models.Child) (*models.ParentWithChildren, error) {
	children := make([]models.Child, len(kid))

	if len(kid) > 0 {
		for i, c := range kid {
			if c.UUID == uuid.Nil {
				c.UUID = uuid.New()
			}
			c.ParentUUID = parent.ID
			child, err := p.childService.PostChildDatum(ctx, c)
			if err != nil {
				return nil, errors.Wrap(err, "failed to add child data")
			}
			children[i] = *child
		}
	}
	parentModel := parent.ToDbModel()
	pDbM, err := p.parentDao.AddParentData(ctx, parentModel)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add parent data")
	}
	pM := models.ParentFromDb(*pDbM)
	return &models.ParentWithChildren{
		Par:      *pM,
		Children: children,
	}, nil
}
