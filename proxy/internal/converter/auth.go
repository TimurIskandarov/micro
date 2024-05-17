package converter

import (
	"proxy/internal/model"

	proto "auth/pkg/auth_v1"
)

func ToRegisterRequestProto(req *model.RegisterIn) *proto.RegisterRequest {
	return &proto.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToRegisterOut(resp *proto.RegisterResponse) *model.RegisterOut {
	return &model.RegisterOut{
		Message: resp.GetMessage(),
		Status:  int(resp.GetStatus()),
	}
}

func ToLoginRequestProto(req *model.AuthorizedIn) *proto.LoginRequest {
	return &proto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToAuthorizedOut(resp *proto.LoginResponse) *model.AuthorizedOut {
	return &model.AuthorizedOut{
		Name:        resp.GetMessage(),
		AccessToken: resp.GetToken(),
		Message:     resp.GetMessage(),
		Status:      int(resp.GetStatus()),
	}
}
