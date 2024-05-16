package auth_v1

import (
	"context"

	proto "auth/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// out, err := i.authService.Login(ctx, *converter.ToLoginRequest(req))
	// if err != nil {
	// 	return nil, err
	// }

	// return converter.ToLoginResponseProto(out), nil
	return &proto.LoginResponse{}, nil
}
