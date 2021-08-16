package eth1

import (
	"context"
	"testing"

	"github.com/consensys/quorum-key-manager/pkg/errors"
	"github.com/consensys/quorum-key-manager/src/infra/log/testutils"
	"github.com/consensys/quorum-key-manager/src/stores/database"
	mock2 "github.com/consensys/quorum-key-manager/src/stores/database/mock"
	testutils2 "github.com/consensys/quorum-key-manager/src/stores/entities/testutils"
	"github.com/consensys/quorum-key-manager/src/stores/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteKey(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock.NewMockKeyStore(ctrl)
	db := mock2.NewMockETH1Accounts(ctrl)
	logger := testutils.NewMockLogger(ctrl)

	connector := NewConnector(store, db, nil, logger)

	db.EXPECT().RunInTransaction(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, persist func(dbtx database.ETH1Accounts) error) error {
			return persist(db)
		}).AnyTimes()

	t.Run("should delete eth1Account successfully", func(t *testing.T) {
		acc := testutils2.FakeETH1Account()
		key := testutils2.FakeKey()
		key.ID = acc.KeyID

		db.EXPECT().Get(gomock.Any(), acc.Address.Hex()).Return(acc, nil)

		db.EXPECT().Delete(gomock.Any(), acc.Address.Hex()).Return(nil)

		store.EXPECT().Delete(gomock.Any(), key.ID).Return(nil)

		err := connector.Delete(ctx, acc.Address)

		assert.NoError(t, err)
	})

	t.Run("should delete key successfully, ignoring not supported error", func(t *testing.T) {
		rErr := errors.NotSupportedError("not supported")
		acc := testutils2.FakeETH1Account()
		key := testutils2.FakeKey()
		key.ID = acc.KeyID

		db.EXPECT().Get(gomock.Any(), acc.Address.Hex()).Return(acc, nil)

		db.EXPECT().Delete(gomock.Any(), acc.Address.Hex()).Return(nil)

		store.EXPECT().Delete(gomock.Any(), key.ID).Return(rErr)

		err := connector.Delete(ctx, acc.Address)

		assert.NoError(t, err)
	})

	t.Run("should fail to delete key if db fail to get", func(t *testing.T) {
		expectedErr := errors.NotFoundError("not found")
		acc := testutils2.FakeETH1Account()
		key := testutils2.FakeKey()
		key.ID = acc.KeyID

		db.EXPECT().Get(gomock.Any(), acc.Address.Hex()).Return(acc, expectedErr)

		err := connector.Delete(ctx, acc.Address)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("should fail to delete key if db fail to delete", func(t *testing.T) {
		expectedErr := errors.PostgresError("cannot connect")
		acc := testutils2.FakeETH1Account()
		key := testutils2.FakeKey()
		key.ID = acc.KeyID

		db.EXPECT().Get(gomock.Any(), acc.Address.Hex()).Return(acc, nil)

		db.EXPECT().Delete(gomock.Any(), acc.Address.Hex()).Return(expectedErr)

		err := connector.Delete(ctx, acc.Address)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("should fail to delete key if store fail to delete", func(t *testing.T) {
		expectedErr := errors.UnauthorizedError("not authorized")
		acc := testutils2.FakeETH1Account()
		key := testutils2.FakeKey()
		key.ID = acc.KeyID

		db.EXPECT().Get(gomock.Any(), acc.Address.Hex()).Return(acc, nil)

		db.EXPECT().Delete(gomock.Any(), acc.Address.Hex()).Return(nil)

		store.EXPECT().Delete(gomock.Any(), key.ID).Return(expectedErr)

		err := connector.Delete(ctx, acc.Address)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})
}
