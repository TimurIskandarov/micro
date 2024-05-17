package auth_v1

import (
	"context"

	"auth/internal/converter"
	proto "auth/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	out := i.authService.Login(ctx, *converter.ToLoginRequest(req))

	return converter.ToLoginResponseProto(out), nil
}
