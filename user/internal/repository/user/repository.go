package user

import (
	"context"

	"user/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	Create(ctx context.Context, user *model.User) error
	Profile(ctx context.Context, email string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
}

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}
