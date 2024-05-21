package user

import (
	"context"

	"user/internal/model"
	// "go.uber.org/zap"
)

func (s *service) Create(ctx context.Context, in *model.User) (int64, error) {
	// out, err := s.repo.CallExternalApi(ctx, "geolocate", in.Lat, in.Lng)
	// if err != nil {
	// 	s.logger.Error("search geocode failed", zap.Error(err))
	// 	return model.GeocodeOut{}, err
	// }

	return 0, nil
}
