package geo

import (
	"context"

	"proxy/internal/converter"
	"proxy/internal/model"
)

func (s *service) Geocode(ctx context.Context, in *model.GeocodeIn) (*model.GeocodeOut, error) {
	resp, err := s.client.GeoCode(ctx, converter.ToGeocodeRequestProto(in))
	if err != nil {
		return nil, err
	}

	return converter.ToGeocodeOut(resp), nil
}
