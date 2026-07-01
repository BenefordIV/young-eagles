package dao

import (
	"context"
	"log"
	"young-eagles/external/models"
	dbmodels "young-eagles/internal/db/gen/models"

	"github.com/aarondl/opt/omit"
	"github.com/gofrs/uuid/v5"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

type ChildrenDao interface {
	AddChildData(ctx context.Context, child models.Child) (*models.Child, error)
	GetChildByFirstLastName(ctx context.Context, fn, ln string) (*models.Child, error)
	GetChildByUUID(ctx context.Context, id uuid.UUID) (*models.Child, error)
}

type childrenDaoImpl struct {
	dbConn bob.DB
}

func NewChildrenDao(conn bob.DB) ChildrenDao {
	return &childrenDaoImpl{
		dbConn: conn,
	}
}

func (c childrenDaoImpl) AddChildData(ctx context.Context, child models.Child) (*models.Child, error) {
	s := &dbmodels.ChildSetter{
		FirstName:      omit.From(child.FirstName),
		LastName:       omit.From(child.LastName),
		DateOfBirth:    omit.From(child.DateOfBirth),
		HasCertificate: omit.From(child.HasCertificate),
	}
	ch, err := dbmodels.Children.Insert(s).One(ctx, c.dbConn)
	if err != nil {
		return nil, err
	}

	return models.ChildFromDb(*ch), nil
}

func (c childrenDaoImpl) GetChildByFirstLastName(ctx context.Context, fn string, ln string) (*models.Child, error) {
	log.Printf("getting child by firstname/lastname %s, %s", ln, fn)
	s := []bob.Mod[*dialect.SelectQuery]{
		dbmodels.SelectWhere.Children.FirstName.EQ(fn),
		dbmodels.SelectWhere.Children.LastName.EQ(ln),
	}

	row, err := dbmodels.Children.Query(s...).One(ctx, c.dbConn)
	if err != nil {
		return nil, err
	}

	return models.ChildFromDb(*row), nil
}

func (c childrenDaoImpl) GetChildByUUID(ctx context.Context, id uuid.UUID) (*models.Child, error) {
	log.Printf("getting child by uuid %v", id)
	row, err := dbmodels.Children.Query(dbmodels.SelectWhere.Children.ID.EQ(id)).One(ctx, c.dbConn)
	if err != nil {
		return nil, err
	}

	return models.ChildFromDb(*row), nil
}
