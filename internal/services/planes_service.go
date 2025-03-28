package services

import (
	"context"
	"errors"
	"github.com/volatiletech/null/v8"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
	"young-eagles/internal/dbmodels"
)

type PlanesService interface {
	AddPlaneDatum(ctx context.Context, callNumber, planeModel, planeMake string) (*models.Plane, error)
	DeletePlaneDatum(ctx context.Context, number string) error
	ReinstatePlaneDatum(ctx context.Context, number string) (*models.Plane, error)
}

type planesServiceImpl struct {
	planeDao dao.PlaneDao
}

func MakePlanesService(planeDao dao.PlaneDao) PlanesService {
	return &planesServiceImpl{
		planeDao: planeDao,
	}
}

func (p planesServiceImpl) AddPlaneDatum(ctx context.Context, callNumber, planeModel, planeMake string) (*models.Plane, error) {
	plane, _ := p.planeDao.FindPlaneByCallNumber(ctx, callNumber)

	if plane != nil {
		return nil, errors.New("plane already exists, cannot add")
	}

	// Add plane to database
	pAdd := dbmodels.PlaneInformation{
		CallNumber: callNumber,
		Model:      null.StringFrom(planeModel),
		Make:       null.StringFrom(planeMake),
	}

	pAdded, err := p.planeDao.AddPlaneDatum(ctx, pAdd)
	if err != nil {
		return nil, err
	}

	pM := models.PlaneFromDb(*pAdded)

	return &pM, nil
}

func (p planesServiceImpl) DeletePlaneDatum(ctx context.Context, number string) error {
	plane, err := p.planeDao.FindPlaneByCallNumber(ctx, number)
	if err != nil {
		return err
	}

	if plane == nil {
		return errors.New("plane not found")
	}

	err = p.planeDao.DeletePlane(ctx, plane)

	if err != nil {
		return nil
	}

	return nil
}

func (p planesServiceImpl) ReinstatePlaneDatum(ctx context.Context, number string) (*models.Plane, error) {
	plane, err := p.planeDao.FindPlaneByCallNumber(ctx, number)
	if err != nil {
		return nil, err
	}

	if plane == nil {
		return nil, errors.New("plane not found")
	}

	plane.DeletedTS = null.Time{}
	// Set DeletedTS to null to reinstate the plane
	err = p.planeDao.UpdatePlane(ctx, plane)

	// Reinstate the plane (this might involve updating a status or similar)
	// For now, we will just return the plane as is
	pM := models.PlaneFromDb(*plane)

	return &pM, nil
}
