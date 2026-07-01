package dao_test

import (
	"context"
	"crypto/rand"
	"testing"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
	"young-eagles/internal/db/gen/models/factory"
	"young-eagles/test"

	"github.com/stephenafamo/bob"
	"github.com/stretchr/testify/require"
)

func Test_GetPlaneByCallNumber(t *testing.T) {
	db := test.ConnectDB(t)
	bdb := bob.NewDB(db)
	defer bdb.Close()
	planeDao := dao.NewPlaneDao(bdb)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		f := factory.New()
		tp, err := f.NewPlane().Create(ctx, bdb)
		require.NoError(t, err)

		fp, err := planeDao.FindPlaneByCallNumber(ctx, tp.CallNumber)
		require.NoError(t, err)
		require.Equal(t, tp.CallNumber, fp.CallNumber)
		require.Equal(t, tp.Model, fp.PlaneModel)
		require.Equal(t, tp.Make, fp.PlaneMake)
	})

	t.Run("not found", func(t *testing.T) {
		f := factory.New()
		_, err := f.NewPlane().Create(ctx, bdb)
		require.NoError(t, err)
		_, err = planeDao.FindPlaneByCallNumber(ctx, "TEST")
		require.Error(t, err)
	})
}

func Test_AddNewPlane(t *testing.T) {
	db := test.ConnectDB(t)
	bdb := bob.NewDB(db)
	defer bdb.Close()
	planeDao := dao.NewPlaneDao(bdb)
	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		p := models.Plane{
			CallNumber: rand.Text(),
			PlaneModel: "Cesna",
			PlaneMake:  "Skyhawk",
		}

		added, err := planeDao.AddPlaneDatum(ctx, p)
		require.NoError(t, err)
		require.Equal(t, p.CallNumber, added.CallNumber)
		require.Equal(t, p.PlaneModel, added.PlaneModel)
		require.Equal(t, p.PlaneMake, added.PlaneMake)

		found, err := planeDao.FindPlaneByCallNumber(ctx, p.CallNumber)
		require.NoError(t, err)
		require.Equal(t, p.CallNumber, found.CallNumber)
	})
}

func Test_DeletePlane(t *testing.T) {
	db := test.ConnectDB(t)
	bdb := bob.NewDB(db)
	defer bdb.Close()
	planeDao := dao.NewPlaneDao(bdb)
	ctx := context.Background()
	f := factory.New()
	tp, err := f.NewPlane().Create(ctx, bdb)
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		err := planeDao.DeletePlane(ctx, &models.Plane{
			CallNumber: tp.CallNumber,
			PlaneModel: tp.Model,
			PlaneMake:  tp.Make,
		})
		require.NoError(t, err)
	})
}

func Test_UpdatePlane(t *testing.T) {
	db := test.ConnectDB(t)
	bdb := bob.NewDB(db)
	defer bdb.Close()
	planeDao := dao.NewPlaneDao(bdb)
	ctx := context.Background()
	f := factory.New()
	tp, err := f.NewPlane().Create(ctx, bdb)
	require.NoError(t, err)
	t.Run("success", func(t *testing.T) {
		p := models.Plane{
			CallNumber: tp.CallNumber,
			PlaneModel: "test",
			PlaneMake:  "up",
		}
		err = planeDao.UpdatePlane(ctx, &p)
		require.NoError(t, err)
		found, err := planeDao.FindPlaneByCallNumber(ctx, p.CallNumber)
		require.NoError(t, err)
		require.Equal(t, p.CallNumber, found.CallNumber)
		require.Equal(t, p.PlaneModel, found.PlaneModel)
		require.Equal(t, p.PlaneMake, found.PlaneMake)
	})
}
