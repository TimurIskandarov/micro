package user_v1

import (
	"user/internal/service/user"
	proto "user/pkg/user_v1"
)

type Implementation struct {
	userService user.Service
	proto.UnimplementedUserV1Server
}

func NewImplementation(userService user.Service) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
