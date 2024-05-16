package auth

import (
	"context"

	"auth/internal/model"

	"github.com/go-chi/jwtauth"
	"go.uber.org/zap"
)

var (
	users = make(map[string]string)
)

var _ Service = (*service)(nil)

type Service interface {
	Register(ctx context.Context, in model.RegisterIn) model.RegisterOut
	Login(ctx context.Context, in model.AuthorizedIn) model.AuthorizedOut
}

type service struct {
	token *jwtauth.JWTAuth
	logger *zap.Logger
}

func NewService(token *jwtauth.JWTAuth, logger *zap.Logger) *service {
	return &service{
		token: token,
		logger: logger,
	}
}
