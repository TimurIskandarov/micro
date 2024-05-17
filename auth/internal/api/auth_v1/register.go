package auth_v1

import (
	"context"

	"auth/internal/converter"
	proto "auth/pkg/auth_v1"
)

func (i *Implementation) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	out := i.authService.Register(ctx, *converter.ToRegisterRequest(req))

	return converter.ToRegisterResponseProto(out), nil
}
