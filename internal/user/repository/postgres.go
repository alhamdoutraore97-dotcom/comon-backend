package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alhamdoutraore97-dotcom/backend/internal/user/domain"
)

type UserPostgresRepo struct {
	db *sql.DB
}

func NewUserPostgresRepo(db *sql.DB) *UserPostgresRepo {
	return &UserPostgresRepo{db: db}
}

// GetByID récupère un utilisateur par son ID

func (r *UserPostgresRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := "SELECT id, name, email, created_at FROM users WHERE id = $1"

	var u domain.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &u, nil
}

// Create insère un nouvel utilisateur

func (r *UserPostgresRepo) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&user.ID)
}

// Update met à jour un utilisateur

func (r *UserPostgresRepo) Update(ctx context.Context, user *domain.User) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	return err
}

// Delete supprime un utilisateur

func (r *UserPostgresRepo) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
