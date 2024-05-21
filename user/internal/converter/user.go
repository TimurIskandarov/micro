package converter

import (
	"user/internal/model"
	proto "user/pkg/user_v1"
)

func ToCreateRequest(req *proto.CreateRequest) *model.User {
	return &model.User{
		Name:     req.GetUser().GetName(),
		Email:    req.GetUser().GetEmail(),
		Phone:    req.GetUser().GetPhone(),
		Password: req.GetUser().GetPassword(),
	}
}

func ToProfileResponseProto(user *model.User) *proto.ProfileResponse {
	return &proto.ProfileResponse{
		User: &proto.User{
			Name:     user.Name,
			Email:    user.Email,
			Phone:    user.Phone,
			Password: user.Password,
		},
		Id: int64(user.ID),
	}
}

func ToListResponseProto(profiles []*model.User) *proto.ListResponse {
	users := make([]*proto.User, len(profiles))

	for i, user := range profiles {
		users[i] = &proto.User{
			Name:     user.Name,
			Email:    user.Email,
			Phone:    user.Phone,
			Password: user.Password,
		}
	}

	return &proto.ListResponse{
		Users: users,
	}
}
