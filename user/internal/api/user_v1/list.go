package user_v1

import (
	"context"
	"user/internal/converter"

	proto "user/pkg/user_v1"
)

func (i *Implementation) List(ctx context.Context) (*proto.ListResponse, error) {
	users, err := i.userService.List(ctx)
	if err != nil {
		return nil, err
	}

	return converter.ToListResponseProto(users), nil
}
