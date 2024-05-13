package geo

import (
	"context"

	"geo/internal/model"

	"go.uber.org/zap"
)

func (s *service) Geocode(ctx context.Context, in model.GeocodeIn) (model.GeocodeOut, error) {
	out, err := s.repo.CallExternalApi(ctx, "geolocate", in.Lat, in.Lng)
	if err != nil {
		s.logger.Error("search geocode failed", zap.Error(err))
		return model.GeocodeOut{}, err
	}

	return model.GeocodeOut{Addresses: out}, nil
}
