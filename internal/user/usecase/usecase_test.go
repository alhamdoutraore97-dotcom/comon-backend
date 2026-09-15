package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/alhamdoutraore97-dotcom/backend/internal/user/domain"
	mock_repo "github.com/alhamdoutraore97-dotcom/backend/internal/user/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============ GetUser ============

func TestGetUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	expected := &domain.User{ID: 1, Name: "Alice", Email: "alice@test.com"}

	mockRepo.EXPECT().GetByID(gomock.Any(), "1").Return(expected, nil)

	uc := NewUserUsecase(mockRepo)
	user, err := uc.GetUser(1)

	require.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
}

func TestGetUser_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), "999").Return(nil, errors.New("not found"))

	uc := NewUserUsecase(mockRepo)
	user, err := uc.GetUser(999)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestGetUser_InvalidID_Zero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	// Aucun appel attendu / Никаких вызовов не ожидается
	uc := NewUserUsecase(mockRepo)

	user, err := uc.GetUser(0)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestGetUser_InvalidID_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	uc := NewUserUsecase(mockRepo)

	user, err := uc.GetUser(-5)
	assert.Error(t, err)
	assert.Nil(t, user)
}

// ============ CreateUser ============

func TestCreateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, u *domain.User) error {
		u.ID = 42
		return nil
	})

	uc := NewUserUsecase(mockRepo)
	user, err := uc.CreateUser("Bob", "bob@test.com")

	require.NoError(t, err)
	assert.Equal(t, 42, user.ID)
	assert.Equal(t, "Bob", user.Name)
}

func TestCreateUser_EmptyName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	uc := NewUserUsecase(mockRepo)

	_, err := uc.CreateUser("", "test@test.com")
	assert.Error(t, err)
}

func TestCreateUser_EmptyEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	uc := NewUserUsecase(mockRepo)

	_, err := uc.CreateUser("Bob", "")
	assert.Error(t, err)
}

func TestCreateUser_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("db error"))

	uc := NewUserUsecase(mockRepo)
	_, err := uc.CreateUser("Bob", "bob@test.com")
	assert.Error(t, err)
}

// ============ UpdateUser (si vous l'ajoutez) ============

func TestUpdateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

	uc := NewUserUsecase(mockRepo)
	err := uc.UpdateUser(&domain.User{ID: 1, Name: "Updated"})
	assert.NoError(t, err)
}

// ============ DeleteUser (si vous l'ajoutez) ============

func TestDeleteUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().Delete(gomock.Any(), "1").Return(nil)

	uc := NewUserUsecase(mockRepo)
	err := uc.DeleteUser(1)
	assert.NoError(t, err)
}
