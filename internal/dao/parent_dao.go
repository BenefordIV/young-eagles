package dao

import (
	"context"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"young-eagles/internal/dbmodels"
)

type ParentDao interface {
	GetParentByEmail(ctx context.Context, email string) (*dbmodels.ParentInformation, error)
	AddParentData(ctx context.Context, model dbmodels.ParentInformation) (*dbmodels.ParentInformation, error)
	GetParentByID(ctx context.Context, id string) (*dbmodels.ParentInformation, error)
}

type parentDaoImpl struct {
	dbConn DbConnection
}

func NewParentDao(conn DbConnection) ParentDao {
	return &parentDaoImpl{
		dbConn: conn,
	}
}

func (p parentDaoImpl) GetParentByEmail(ctx context.Context, email string) (*dbmodels.ParentInformation, error) {
	parent, err := dbmodels.ParentInformations(dbmodels.ParentInformationWhere.Email.EQ(null.StringFrom(email))).One(ctx, p.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return parent, nil
}

func (p parentDaoImpl) AddParentData(ctx context.Context, model dbmodels.ParentInformation) (*dbmodels.ParentInformation, error) {
	err := model.Insert(ctx, p.dbConn.DbConn,
		boil.Blacklist(dbmodels.ParentInformationColumns.UpdatedTS, dbmodels.ParentInformationColumns.DeletedTS))
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (p parentDaoImpl) GetParentByID(ctx context.Context, id string) (*dbmodels.ParentInformation, error) {
	parent, err := dbmodels.ParentInformations(dbmodels.ParentInformationWhere.UUID.EQ(id)).One(ctx, p.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return parent, nil
}
