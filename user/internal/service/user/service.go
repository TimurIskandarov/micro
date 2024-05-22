package user

import (
	"context"

	"user/internal/model"
	"user/internal/repository/user"

	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(ctx context.Context, in *model.User) (int64, error)
	Profile(ctx context.Context, email string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
}

type service struct {
	repo   user.Repository
	logger *zap.Logger
}

func NewService(repo user.Repository, logger *zap.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}
