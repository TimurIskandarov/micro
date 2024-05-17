package converter

import (
	"auth/internal/model"
	proto "auth/pkg/auth_v1"
)

func ToRegisterRequest(req *proto.RegisterRequest) *model.RegisterIn {
	return &model.RegisterIn{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func ToRegisterResponseProto(resp model.RegisterOut) *proto.RegisterResponse {
	return &proto.RegisterResponse{
		Message: resp.Message,
		Status:  uint32(resp.Status),
	}
}

func ToLoginRequest(req *proto.LoginRequest) *model.AuthorizedIn {
	return &model.AuthorizedIn{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func ToLoginResponseProto(resp model.AuthorizedOut) *proto.LoginResponse {
	return &proto.LoginResponse{
		Name: resp.Name,
		Token: resp.AccessToken,
		Message: resp.Message,
		Status:  uint32(resp.Status),
	}
}
