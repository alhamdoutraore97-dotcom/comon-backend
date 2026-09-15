package repository

import (
	"context"
	"strconv"
	"testing"

	"github.com/alhamdoutraore97-dotcom/backend/internal/user/domain"
	"github.com/alhamdoutraore97-dotcom/backend/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupUserRepo initialise le repo pour les tests
// setupUserRepo инициализирует репозиторий для тестов
func setupUserRepo(t *testing.T) *UserPostgresRepo {
	db := testutil.SetupTestDB(t)
	testutil.RunMigrations(t, db)
	return NewUserPostgresRepo(db)
}

// TestUserRepo_Create teste la création d'un utilisateur
// TestUserRepo_Create тестирует создание пользователя
func TestUserRepo_Create(t *testing.T) {
	repo := setupUserRepo(t)

	user := &domain.User{Name: "Charlie", Email: "charlie@test.com"}
	err := repo.Create(context.Background(), user)

	require.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, "Charlie", user.Name)
}

// TestUserRepo_Create_DuplicateEmail teste l'email unique
// TestUserRepo_Create_DuplicateEmail тестирует уникальность email
func TestUserRepo_Create_DuplicateEmail(t *testing.T) {
	repo := setupUserRepo(t)

	user1 := &domain.User{Name: "A", Email: "same@test.com"}
	require.NoError(t, repo.Create(context.Background(), user1))

	user2 := &domain.User{Name: "B", Email: "same@test.com"}
	err := repo.Create(context.Background(), user2)

	assert.Error(t, err) // Violation de contrainte UNIQUE
}

// TestUserRepo_GetByID teste la récupération par ID
// TestUserRepo_GetByID тестирует получение по ID
func TestUserRepo_GetByID(t *testing.T) {
	db := testutil.SetupTestDB(t)
	testutil.RunMigrations(t, db)
	repo := NewUserPostgresRepo(db)

	id := testutil.CreateTestUser(t, db, "David", "david@test.com")

	user, err := repo.GetByID(context.Background(), strconv.Itoa(id))
	require.NoError(t, err)
	assert.Equal(t, id, user.ID)
	assert.Equal(t, "David", user.Name)
}

// TestUserRepo_GetByID_NotFound teste le cas non trouvé
// TestUserRepo_GetByID_NotFound тестирует случай "не найдено"
func TestUserRepo_GetByID_NotFound(t *testing.T) {
	repo := setupUserRepo(t)

	user, err := repo.GetByID(context.Background(), "9999")

	assert.Error(t, err)
	assert.Nil(t, user)
}

// TestUserRepo_Update teste la mise à jour
// TestUserRepo_Update тестирует обновление
func TestUserRepo_Update(t *testing.T) {
	db := testutil.SetupTestDB(t)
	testutil.RunMigrations(t, db)
	repo := NewUserPostgresRepo(db)

	id := testutil.CreateTestUser(t, db, "Old", "old@test.com")

	user := &domain.User{ID: id, Name: "New", Email: "new@test.com"}
	err := repo.Update(context.Background(), user)
	require.NoError(t, err)

	updated, err := repo.GetByID(context.Background(), strconv.Itoa(id))
	require.NoError(t, err)
	assert.Equal(t, "New", updated.Name)
}

// TestUserRepo_Delete teste la suppression
// TestUserRepo_Delete тестирует удаление
func TestUserRepo_Delete(t *testing.T) {
	db := testutil.SetupTestDB(t)
	testutil.RunMigrations(t, db)
	repo := NewUserPostgresRepo(db)

	id := 0
	err := repo.Delete(context.Background(), strconv.Itoa(id))
	require.NoError(t, err)

	_, err = repo.GetByID(context.Background(), strconv.Itoa(id))
	assert.Error(t, err)
}

// TestUserRepo_Delete_NonExistent teste la suppression inexistante
// TestUserRepo_Delete_NonExistent тестирует удаление несуществующего
func TestUserRepo_Delete_NonExistent(t *testing.T) {
	repo := setupUserRepo(t)
	err := repo.Delete(context.Background(), "9999")
	assert.NoError(t, err) // DELETE ne renvoie pas d'erreur si rien trouvé
}
