package geo

import (
	"context"

	"proxy/internal/converter"
	"proxy/internal/model"
)

func (s *service) AddressSearch(ctx context.Context, in *model.SearchIn) (*model.SearchOut, error) {
	resp, err := s.client.AddressSearch(ctx, converter.ToSearchRequestProto(in))
	if err != nil {
		return nil, err
	}

	return converter.ToSearchOut(resp), nil
}
