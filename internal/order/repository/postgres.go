package repository

import (
	"context"

	"database/sql"
	"fmt"

	"github.com/alhamdoutraore97-dotcom/backend/internal/order/domain"
)

type OrderPostgresRepo struct {
	db *sql.DB
}

func NewOrderPostgresRepo(db *sql.DB) *OrderPostgresRepo {
	return &OrderPostgresRepo{db: db}
}

// GetByID récupère une commande par son ID
// GetByID получает заказ по ID
func (r *OrderPostgresRepo) GetByID(ctx context.Context, id string) (*domain.Order, error) {
	var o domain.Order
	query := "SELECT id, product, quantity FROM orders WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&o.ID, &o.UserID, &o.Product, &o.Quantity, &o.Status)
	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}
	return &o, nil
}

// Create insère une nouvelle commande
// Create вставляет новый заказ
func (r *OrderPostgresRepo) Create(ctx context.Context, order *domain.Order) error {
	query := "INSERT INTO orders (user_id, product, quantity) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, order.UserID, order.Product, order.Quantity).Scan(&order.ID)
}

// UpdateStatus met à jour le statut d'une commande
// UpdateStatus обновляет статус заказа
func (r *OrderPostgresRepo) UpdateStatus(ctx context.Context, id string, status string) error {
	_, err := r.db.Exec("UPDATE orders SET status = $1 WHERE id = $2", status, id)
	return err
}

func (r *OrderPostgresRepo) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *OrderPostgresRepo) Update(ctx context.Context, order *domain.Order) error {
	query := "UPDATE orders SET product = $1, quantity = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, order.Product, order.Quantity, order.ID)
	return err
}
