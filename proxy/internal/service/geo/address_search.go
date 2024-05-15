package geo

import (
	"context"

	"proxy/internal/converter"
	"proxy/internal/model"
)

func (s *service) AddressSearch(ctx context.Context, in model.SearchIn) (*model.SearchOut, error) {
	resp, err := s.client.AddressSearch(ctx, converter.ToSearchRequestProto(in))
	if err != nil {
		return nil, err
	}

	var addresses []*model.Address
	for _, address := range resp.Addresses {
		addresses = append(addresses, &model.Address{
			Lat:   address.Lat,
			Lng:   address.Lon,
			Value: address.Value,
		})
	}

	return &model.SearchOut{
		Addresses: addresses,
	}, nil
}
