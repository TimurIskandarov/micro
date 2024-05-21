package user_v1

import (
	"context"
	"user/internal/converter"

	proto "user/pkg/user_v1"
)

func (i *Implementation) Profile(ctx context.Context, req *proto.ProfileRequest) (*proto.ProfileResponse, error) {
	out, err := i.userService.Profile(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}

	return converter.ToProfileResponseProto(out), nil
}
