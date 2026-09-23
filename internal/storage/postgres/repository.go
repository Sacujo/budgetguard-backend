package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sacujo/budgetguard-backend/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const uniqueViolation = "23505"

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

func (r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	err := r.pool.QueryRow(ctx,
		`INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id, revision_ctr, created_at, updated_at`,
		u.Email,
		u.PasswordHash,
	).Scan(
		&u.ID,
		&u.RevisionCtr,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
		return domain.ErrAlreadyExists
	}
	if err != nil {
		return fmt.Errorf("создание пользователя: %w", err)
	}

	return nil
}
