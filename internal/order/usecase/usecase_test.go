package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/alhamdoutraore97-dotcom/backend/internal/order/domain"
	mock_repo "github.com/alhamdoutraore97-dotcom/backend/internal/order/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), "1").Return(&domain.Order{ID: 1, Product: "Book"}, nil)

	uc := NewOrderUsecase(mockRepo)
	o, err := uc.GetOrder(1)

	require.NoError(t, err)
	assert.Equal(t, "Book", o.Product)
}

func TestGetOrder_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	uc := NewOrderUsecase(mockRepo)

	_, err := uc.GetOrder(0)
	assert.Error(t, err)
}

func TestGetOrder_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), "5").Return(nil, errors.New("not found"))

	uc := NewOrderUsecase(mockRepo)
	o, err := uc.GetOrder(5)

	assert.Error(t, err)
	assert.Nil(t, o)
}

func TestCreateOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, o *domain.Order) error {
		o.ID = 10
		return nil
	})

	uc := NewOrderUsecase(mockRepo)
	o, err := uc.CreateOrder(1, "Livre", 2)

	require.NoError(t, err)
	assert.Equal(t, 10, o.ID)
	assert.Equal(t, "Livre", o.Product)
}

func TestCreateOrder_InvalidUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	uc := NewOrderUsecase(mockRepo)

	_, err := uc.CreateOrder(0, "X", 1)
	assert.Error(t, err)
}

func TestCreateOrder_EmptyProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	uc := NewOrderUsecase(mockRepo)

	_, err := uc.CreateOrder(1, "", 1)
	assert.Error(t, err)
}

func TestCreateOrder_InvalidQuantity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repo.NewMockOrderRepository(ctrl)
	uc := NewOrderUsecase(mockRepo)

	_, err := uc.CreateOrder(1, "X", 0)
	assert.Error(t, err)
}
