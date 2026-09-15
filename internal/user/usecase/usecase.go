package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/alhamdoutraore97-dotcom/backend/internal/user/domain"
	"github.com/alhamdoutraore97-dotcom/backend/internal/user/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// GetUser récupère un utilisateur avec validation
// GetUser получает пользователя с валидацией
func (u *UserUsecase) GetUser(id int) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}
	return u.repo.GetByID(context.Background(), strconv.Itoa(id))
}

// CreateUser crée un utilisateur avec validation
// CreateUser создаёт пользователя с валидацией
func (u *UserUsecase) CreateUser(name, email string) (*domain.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email required")
	}
	user := &domain.User{Name: name, Email: email}
	err := u.repo.Create(context.Background(), user)
	return user, err
}

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	if user == nil || user.ID <= 0 {
		return errors.New("invalid user")
	}
	return u.repo.Update(context.Background(), user)
}

func (u *UserUsecase) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return u.repo.Delete(context.Background(), strconv.Itoa(id))
}
