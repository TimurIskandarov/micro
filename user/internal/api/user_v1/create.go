package user_v1

import (
	"context"
	
	"user/internal/converter"

	proto "user/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Id: id,
	}, nil
}
