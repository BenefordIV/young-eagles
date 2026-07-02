package dao

import (
	"context"
	"young-eagles/external/models"
	dbmodels "young-eagles/internal/db/gen/models"

	"github.com/aarondl/opt/omit"
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
	plane, err := dbmodels.Planes.Query(dbmodels.SelectWhere.Planes.CallNumber.EQ(number)).One(ctx, p.dbConn)
	if err != nil {
		return nil, err
	}

	found := models.PlaneFromDb(*plane)

	return found, nil
}
func (p planeDaoImpl) AddPlaneDatum(ctx context.Context, information models.Plane) (*models.Plane, error) {
	s := &dbmodels.PlaneSetter{
		CallNumber: omit.From(information.CallNumber),
		Model:      omit.From(information.PlaneModel),
		Make:       omit.From(information.PlaneMake),
	}

	row, err := dbmodels.Planes.Insert(s).One(ctx, p.dbConn)
	if err != nil {
		return nil, err
	}

	return models.PlaneFromDb(*row), nil
}

func (p planeDaoImpl) DeletePlane(ctx context.Context, plane *models.Plane) error {
	_, err := dbmodels.Planes.Delete(dbmodels.DeleteWhere.Planes.CallNumber.EQ(plane.CallNumber)).Exec(ctx, p.dbConn)
	if err != nil {
		return err
	}
	return nil
}

func (p planeDaoImpl) UpdatePlane(ctx context.Context, plane *models.Plane) error {
	existing, err := dbmodels.Planes.Query(dbmodels.SelectWhere.Planes.CallNumber.EQ(plane.CallNumber)).One(ctx, p.dbConn)
	if err != nil {
		return err
	}

	s := &dbmodels.PlaneSetter{
		CallNumber: omit.From(plane.CallNumber),
		Model:      omit.From(plane.PlaneModel),
		Make:       omit.From(plane.PlaneMake),
	}

	if err = existing.Update(ctx, p.dbConn, s); err != nil {
		return err
	}
	return nil
}
