package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"young-eagles/internal/dbmodels"
)

type ChildrenDao interface {
	GetChildByID(ctx context.Context, id string) (*dbmodels.ChildInformation, error)
	PostChildData(ctx context.Context, child dbmodels.ChildInformation) (*dbmodels.ChildInformation, error)
	GetChildrenByParentID(ctx context.Context, parentUUID uuid.UUID) (dbmodels.ChildInformationSlice, error)
}

type childrenDaoImpl struct {
	dbConn DbConnection
}

func NewChildrenDao(conn DbConnection) ChildrenDao {
	return &childrenDaoImpl{
		dbConn: conn,
	}
}

func (c childrenDaoImpl) GetChildByID(ctx context.Context, id string) (*dbmodels.ChildInformation, error) {
	panic("haha")
}

func (c childrenDaoImpl) PostChildData(ctx context.Context, child dbmodels.ChildInformation) (*dbmodels.ChildInformation, error) {
	err := child.Insert(ctx, c.dbConn.DbConn,
		boil.Blacklist(dbmodels.ChildInformationColumns.UpdatedTS, dbmodels.ChildInformationColumns.DeletedTS))
	if err != nil {
		return nil, err
	}

	return &child, nil
}

func (c childrenDaoImpl) GetChildrenByParentID(ctx context.Context, parentUUID uuid.UUID) (dbmodels.ChildInformationSlice, error) {
	children, err := dbmodels.ChildInformations(dbmodels.ChildInformationWhere.ParentID.EQ(parentUUID.String())).All(ctx, c.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return children, nil
}
