package geo

import (
	"context"

	"proxy/internal/converter"
	"proxy/internal/model"
)

func (s *service) Geocode(ctx context.Context, in model.GeocodeIn) (*model.GeocodeOut, error) {
	resp, err := s.client.GeoCode(ctx, converter.ToGeocodeRequestProto(in))
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

	return &model.GeocodeOut{
		Addresses: addresses,
	}, nil
}
