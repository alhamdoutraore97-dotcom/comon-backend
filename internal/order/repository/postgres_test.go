package repository

import (
	"context"
	"strconv"
	"testing"

	"github.com/alhamdoutraore97-dotcom/backend/internal/order/domain"
	"github.com/alhamdoutraore97-dotcom/backend/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupOrderRepo(t *testing.T) (*OrderPostgresRepo, int) {
	db := testutil.SetupTestDB(t)
	testutil.RunMigrations(t, db)
	userID := testutil.CreateTestUser(t, db, "OrderUser", "order@test.com")
	return NewOrderPostgresRepo(db), userID
}

func TestOrderRepo_Create(t *testing.T) {
	repo, userID := setupOrderRepo(t)

	order := &domain.Order{UserID: userID, Product: "Livre", Quantity: 2}
	err := repo.Create(context.Background(), order)

	require.NoError(t, err)
	assert.NotZero(t, order.ID)
}

func TestOrderRepo_GetByID(t *testing.T) {
	repo, userID := setupOrderRepo(t)

	order := &domain.Order{UserID: userID, Product: "Stylo", Quantity: 5}
	require.NoError(t, repo.Create(context.Background(), order))

	found, err := repo.GetByID(context.Background(), "order.ID")
	require.NoError(t, err)
	assert.Equal(t, "Stylo", found.Product)
	assert.Equal(t, 5, found.Quantity)
	assert.Equal(t, "pending", found.Status) // valeur par défaut
}

func TestOrderRepo_GetByID_NotFound(t *testing.T) {
	repo, _ := setupOrderRepo(t)

	o, err := repo.GetByID(context.Background(), "9999")
	assert.Error(t, err)
	assert.Nil(t, o)
}

func TestOrderRepo_UpdateStatus(t *testing.T) {
	repo, userID := setupOrderRepo(t)

	order := &domain.Order{UserID: userID, Product: "Cahier", Quantity: 1}
	require.NoError(t, repo.Create(context.Background(), order))

	err := repo.UpdateStatus(context.Background(), strconv.Itoa(order.ID), "shipped")
	require.NoError(t, err)

	updated, err := repo.GetByID(context.Background(), strconv.Itoa(order.ID))
	require.NoError(t, err)
	assert.Equal(t, "shipped", updated.Status)
}

func TestOrderRepo_Create_InvalidUser(t *testing.T) {
	repo, _ := setupOrderRepo(t)

	order := &domain.Order{UserID: 99999, Product: "X", Quantity: 1}
	err := repo.Create(context.Background(), order)
	assert.Error(t, err) // FK violation
}
