package auth

import (
	"context"

	"proxy/internal/model"
	
	proto "auth/pkg/auth_v1"
)

var _ Service = (*service)(nil)

type Service interface {
	Register(ctx context.Context, in *model.RegisterIn) *model.RegisterOut
	Login(ctx context.Context, in *model.AuthorizedIn) *model.AuthorizedOut
}

type service struct {
	client proto.AuthV1Client
}

func NewService(client proto.AuthV1Client) *service {
	return &service{
		client: client,
	}
}
