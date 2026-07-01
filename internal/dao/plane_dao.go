package dao

import (
	"context"
	"young-eagles/internal/db/gen/models"

	"github.com/stephenafamo/bob"
)

type PlaneDao interface {
	AddPlaneDatum(ctx context.Context, information models.Plane) (*models.Plane, error)
	FindPlaneByCallNumber(ctx context.Context, number string) (*models.Plane, error)
	DeletePlane(ctx context.Context, plane *models.Plane) error
	UpdatePlane(ctx context.Context, plane *models.Plane) error
}

type planeDaoImpl struct {
	dbConn bob.DB
}

func NewPlaneDao(conn bob.DB) PlaneDao {
	return &planeDaoImpl{
		dbConn: conn,
	}
}

func (p planeDaoImpl) FindPlaneByCallNumber(ctx context.Context, number string) (*models.Plane, error) {
	//plane, err := dbmodels.PlaneInformations(qm.WithDeleted(), dbmodels.PlaneInformationWhere.CallNumber.EQ(number)).One(ctx, p.dbConn.DbConn)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return plane, nil
	panic("implement me")
}
func (p planeDaoImpl) AddPlaneDatum(ctx context.Context, information models.Plane) (*models.Plane, error) {
	//err := information.Insert(ctx, p.dbConn.DbConn, boil.Blacklist(dbmodels.PlaneInformationColumns.DeletedTS))
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &information, nil
	panic("implement me")
}

func (p planeDaoImpl) DeletePlane(ctx context.Context, plane *models.Plane) error {
	//_, err := plane.Delete(ctx, p.dbConn.DbConn, false)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
	panic("implement me")
}

func (p planeDaoImpl) UpdatePlane(ctx context.Context, plane *models.Plane) error {
	//_, err := plane.Update(ctx, p.dbConn.DbConn, boil.Infer())
	//if err != nil {
	//	return err
	//}
	//
	//return nil
	panic("implement me")
}
