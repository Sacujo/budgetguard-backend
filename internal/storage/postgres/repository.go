package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sacujo/budgetguard-backend/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User

	err := r.pool.QueryRow(ctx,
		`SELECT id, email, password_hash, revision_ctr, created_at, updated_at
	FROM users
	WHERE email = $1`, email).Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.RevisionCtr,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("получение пользователя по email: %w", err)
	}

	return &u, nil
}
