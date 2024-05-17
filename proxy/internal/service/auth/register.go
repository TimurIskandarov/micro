package auth

import (
	"context"

	"proxy/internal/model"
	"proxy/internal/converter"
)

func (s *service) Register(ctx context.Context, in *model.RegisterIn) *model.RegisterOut {
	resp, err := s.client.Register(ctx, converter.ToRegisterRequestProto(in))
	if err != nil {
		return &model.RegisterOut{
			Status: 1,
			Message: err.Error(),
		}
	}

	return converter.ToRegisterOut(resp)
}
