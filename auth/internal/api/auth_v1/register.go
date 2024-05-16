package auth_v1

import (
	"context"

	proto "auth/pkg/auth_v1"
)

func (i *Implementation) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	// out, err := i.authService.Register(ctx, *converter.ToRegisterRequest(req))
	// if err != nil {
	// 	return nil, err
	// }

	// return converter.ToSearchResponseProto(out), nil
	return &proto.RegisterResponse{}, nil
}
