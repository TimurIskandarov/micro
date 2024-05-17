package auth

import (
	"context"

	"proxy/internal/model"
	"proxy/internal/converter"
)

func (s *service) Login(ctx context.Context, in *model.AuthorizedIn) *model.AuthorizedOut {
	resp, err := s.client.Login(ctx, converter.ToLoginRequestProto(in))
	if err != nil {
		return &model.AuthorizedOut{
			Status: 1,
			Message: err.Error(),
		}
	}

	return converter.ToAuthorizedOut(resp)
}
