package dao_test

import (
	"context"
	"testing"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
	"young-eagles/internal/db/gen/models/factory"
	"young-eagles/test"

	"github.com/stephenafamo/bob"
	"github.com/stretchr/testify/require"
)

func TestAddPilotDao(t *testing.T) {
	db := test.ConnectDB(t)
	bdb := bob.NewDB(db)
	defer bdb.Close()

	pilotDao := dao.NewPilotDao(bdb)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		p := models.Pilot{
			PilotFirstName:   "Jim",
			PilotLastName:    "Bob",
			PilotEmail:       "jbob@test.com",
			EaaChapterNumber: 123,
		}

		added, err := pilotDao.AddPilot(ctx, p)
		require.NoError(t, err)
		require.NotEmpty(t, added)
	})
}

func TestGetPilotByNameChapterCombo(t *testing.T) {
	db := test.ConnectDB(t)
	bdb := bob.NewDB(db)
	defer bdb.Close()
	pilotDao := dao.NewPilotDao(bdb)
	ctx := context.Background()

	f := factory.New()
	testPilot, err := f.NewPilot(
		factory.PilotMods.FirstName("James"),
		factory.PilotMods.LastName("Bob"),
		factory.PilotMods.EaaChapterFunc(func() int32 { return 123 }),
	).Create(ctx, bdb)
	require.NoError(t, err)
	require.NotEmpty(t, testPilot)
	t.Run("success", func(t *testing.T) {
		foundPilot, err := pilotDao.GetPilotByNameChapterCombo(ctx,
			testPilot.FirstName,
			testPilot.LastName,
			int(testPilot.EaaChapter))
		require.NoError(t, err)
		require.NotNil(t, foundPilot)
		require.Equal(t, testPilot.FirstName, foundPilot.PilotFirstName)
		require.Equal(t, testPilot.LastName, foundPilot.PilotLastName)
	})
}

func Test_GetPilotByUUID(t *testing.T) {
	db := test.ConnectDB(t)
	bdb := bob.NewDB(db)
	defer bdb.Close()
	pilotDao := dao.NewPilotDao(bdb)
	ctx := context.Background()
	f := factory.New()
	tp, err := f.NewPilot(
		factory.PilotMods.FirstName("James"),
		factory.PilotMods.LastName("Bob"),
		factory.PilotMods.EaaChapterFunc(func() int32 { return 123 }),
	).Create(ctx, bdb)
	require.NoError(t, err)
	require.NotEmpty(t, tp)
	t.Run("success", func(t *testing.T) {
		found, err := pilotDao.GetPilotByUUID(ctx, tp.ID)
		require.NoError(t, err)
		require.NotNil(t, found)
		require.Equal(t, tp.ID, found.PilotUuid)
		require.Equal(t, tp.FirstName, found.PilotFirstName)
		require.Equal(t, tp.LastName, found.PilotLastName)
	})
}

func TestUpdatePilot(t *testing.T) {
	db := test.ConnectDB(t)
	bdb := bob.NewDB(db)
	defer bdb.Close()
	pilotDao := dao.NewPilotDao(bdb)
	ctx := context.Background()
	f := factory.New()
	tp, err := f.NewPilot(
		factory.PilotMods.FirstName("Test"),
		factory.PilotMods.LastName("Pilot"),
		factory.PilotMods.Email("testemail@test.com"),
		factory.PilotMods.EaaChapterFunc(func() int32 { return 123 }),
	).Create(ctx, bdb)
	require.NoError(t, err)
	require.NotEmpty(t, tp)
	t.Run("success", func(t *testing.T) {
		update := models.Pilot{
			PilotUuid:        tp.ID,
			PilotFirstName:   "New",
			PilotLastName:    "Name",
			PilotEmail:       "newemail@test.com",
			EaaChapterNumber: 456,
		}

		updated, err := pilotDao.UpdatePilot(ctx, &update)
		require.NoError(t, err)
		require.NotNil(t, updated)
		require.Equal(t, update.PilotFirstName, updated.PilotFirstName)
		require.Equal(t, update.PilotLastName, updated.PilotLastName)
		require.Equal(t, update.EaaChapterNumber, updated.EaaChapterNumber)
		require.Equal(t, update.PilotEmail, updated.PilotEmail)
	})
}
