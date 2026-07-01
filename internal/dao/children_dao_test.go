package dao_test

import (
	"context"
	"testing"
	"time"
	"young-eagles/external/models"
	"young-eagles/internal/dao"
	"young-eagles/internal/db/gen/models/factory"
	"young-eagles/test"

	"github.com/stephenafamo/bob"
	"github.com/stretchr/testify/require"
)

func TestChildrenAdd(t *testing.T) {
	cDb := test.ConnectDB(t)
	bdb := bob.NewDB(cDb)
	defer bdb.Close()

	childDao := dao.NewChildrenDao(bdb)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {

		c := models.Child{
			FirstName:      "Test",
			LastName:       "Child",
			DateOfBirth:    time.Now(),
			HasCertificate: false,
		}

		added, err := childDao.AddChildData(ctx, c)
		require.NoError(t, err)
		require.NotEmpty(t, added)

		got, err := childDao.GetChildByUUID(ctx, added.UUID)
		require.NoError(t, err)
		require.Equal(t, got.UUID, got.UUID)
	})
}

func TestChildren_GetByFirstNameLastName(t *testing.T) {
	cDb := test.ConnectDB(t)
	bdb := bob.NewDB(cDb)
	defer bdb.Close()
	childDao := dao.NewChildrenDao(bdb)
	ctx := context.Background()
	f := factory.New()
	testChild, err := f.NewChild(
		factory.ChildMods.FirstName("Bob"),
		factory.ChildMods.LastName("Hope"),
	).Create(ctx, bdb)
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		got, err := childDao.GetChildByFirstLastName(ctx, "Bob", "Hope")
		require.NoError(t, err)
		require.Equal(t, testChild.ID.String(), got.UUID.String())
	})
}
