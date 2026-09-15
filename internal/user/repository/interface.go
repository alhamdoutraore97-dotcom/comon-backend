package repository

import (
	"context"

	"github.com/alhamdoutraore97-dotcom/backend/internal/user/domain"
)

// UserRepository est l'interface que vos implémentations (postgres) et vos mocks doivent respecter.
type UserRepository interface {
	// Ajoutez ici les méthodes que votre usecase utilise.
	// Par exemple :
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}
