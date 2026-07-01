package services

import (
	"context"
	"errors"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
)

type PlanesService interface {
	AddPlaneDatum(ctx context.Context, plane models.Plane) (*models.Plane, error)
	DeletePlaneDatum(ctx context.Context, number string) error
}

type planesServiceImpl struct {
	planeDao dao.PlaneDao
}

func MakePlanesService(planeDao dao.PlaneDao) PlanesService {
	return &planesServiceImpl{
		planeDao: planeDao,
	}
}

func (p planesServiceImpl) AddPlaneDatum(ctx context.Context, plane models.Plane) (*models.Plane, error) {
	existingPlane, _ := p.planeDao.FindPlaneByCallNumber(ctx, plane.CallNumber)

	if existingPlane != nil {
		return nil, errors.New("plane already exists, cannot add")
	}

	// Add plane to database

	pAdded, err := p.planeDao.AddPlaneDatum(ctx, plane)
	if err != nil {
		return nil, err
	}

	return pAdded, nil
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
