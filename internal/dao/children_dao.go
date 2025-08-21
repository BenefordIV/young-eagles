package dao

import (
	"context"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"young-eagles/internal/dbmodels"
)

type ChildrenDao interface {
	PostChildData(ctx context.Context, child dbmodels.ChildInformation) (*dbmodels.ChildInformation, error)
	GetChildByFirstLastName(ctx context.Context, fn, ln string) (*dbmodels.ChildInformation, error)
	GetChildByUUID(ctx context.Context, uuid string) (*dbmodels.ChildInformation, error)
}

type childrenDaoImpl struct {
	dbConn DbConnection
}

func NewChildrenDao(conn DbConnection) ChildrenDao {
	return &childrenDaoImpl{
		dbConn: conn,
	}
}

func (c childrenDaoImpl) PostChildData(ctx context.Context, child dbmodels.ChildInformation) (*dbmodels.ChildInformation, error) {
	log.Println("adding child to database")
	log.Println(child.UUID)
	err := child.Insert(ctx, c.dbConn.DbConn,
		boil.Blacklist(dbmodels.ChildInformationColumns.UpdatedTS, dbmodels.ChildInformationColumns.DeletedTS))
	if err != nil {
		log.Printf("error %v", err)
		return nil, err
	}

	return &child, nil
}

func (c childrenDaoImpl) GetChildByFirstLastName(ctx context.Context, fn string, ln string) (*dbmodels.ChildInformation, error) {
	ec, err := dbmodels.ChildInformations(
		dbmodels.ChildInformationWhere.FirstName.EQ(null.StringFrom(fn)),
		dbmodels.ChildInformationWhere.LastName.EQ(null.StringFrom(ln)),
	).One(ctx, c.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return ec, nil
}

func (c childrenDaoImpl) GetChildByUUID(ctx context.Context, uuid string) (*dbmodels.ChildInformation, error) {
	log.Printf("getting child by uuid %v", uuid)
	ec, err := dbmodels.ChildInformations(
		dbmodels.ChildInformationWhere.UUID.EQ(uuid),
	).One(ctx, c.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return ec, nil
}
