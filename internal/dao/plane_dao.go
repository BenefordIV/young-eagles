package dao

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"young-eagles/internal/dbmodels"
)

type PlaneDao interface {
	AddPlaneDatum(ctx context.Context, information dbmodels.PlaneInformation) (*dbmodels.PlaneInformation, error)
	FindPlaneByCallNumber(ctx context.Context, number string) (*dbmodels.PlaneInformation, error)
	DeletePlane(ctx context.Context, plane *dbmodels.PlaneInformation) error
	UpdatePlane(ctx context.Context, plane *dbmodels.PlaneInformation) error
}

type planeDaoImpl struct {
	dbConn DbConnection
}

func NewPlaneDao(conn DbConnection) PlaneDao {
	return &planeDaoImpl{
		dbConn: conn,
	}
}

func (p planeDaoImpl) FindPlaneByCallNumber(ctx context.Context, number string) (*dbmodels.PlaneInformation, error) {
	plane, err := dbmodels.PlaneInformations(qm.WithDeleted(), dbmodels.PlaneInformationWhere.CallNumber.EQ(number)).One(ctx, p.dbConn.DbConn)
	if err != nil {
		return nil, err
	}

	return plane, nil
}
func (p planeDaoImpl) AddPlaneDatum(ctx context.Context, information dbmodels.PlaneInformation) (*dbmodels.PlaneInformation, error) {
	err := information.Insert(ctx, p.dbConn.DbConn, boil.Blacklist(dbmodels.PlaneInformationColumns.DeletedTS))
	if err != nil {
		return nil, err
	}

	return &information, nil
}

func (p planeDaoImpl) DeletePlane(ctx context.Context, plane *dbmodels.PlaneInformation) error {
	_, err := plane.Delete(ctx, p.dbConn.DbConn, false)
	if err != nil {
		return err
	}

	return nil
}

func (p planeDaoImpl) UpdatePlane(ctx context.Context, plane *dbmodels.PlaneInformation) error {
	_, err := plane.Update(ctx, p.dbConn.DbConn, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
