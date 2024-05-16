package auth_v1

import (
	"auth/internal/service/auth"
	proto "auth/pkg/auth_v1"
)

type Implementation struct {
	authService auth.Service
	proto.UnimplementedAuthV1Server
}

func NewImplementation(authService auth.Service) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
