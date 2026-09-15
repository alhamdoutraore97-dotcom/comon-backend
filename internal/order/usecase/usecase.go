package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/alhamdoutraore97-dotcom/backend/internal/order/domain"
	"github.com/alhamdoutraore97-dotcom/backend/internal/order/repository"
)

type OrderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUsecase(repo repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

// GetOrder récupère une commande avec validation
// GetOrder получает заказ с валидацией
func (u *OrderUsecase) GetOrder(id int) (*domain.Order, error) {
	if id <= 0 {
		return nil, errors.New("invalid order id")
	}
	return u.repo.GetByID(context.Background(), strconv.Itoa(id))
}

// CreateOrder crée une commande avec validation
// CreateOrder создаёт заказ с валидацией
func (u *OrderUsecase) CreateOrder(userID int, product string, quantity int) (*domain.Order, error) {
	if userID <= 0 || product == "" || quantity <= 0 {
		return nil, errors.New("invalid order data")
	}
	order := &domain.Order{UserID: userID, Product: product, Quantity: quantity}
	err := u.repo.Create(context.Background(), order)
	return order, err
}
