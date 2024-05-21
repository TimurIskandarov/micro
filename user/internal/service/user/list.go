package user

import (
	"context"

	"user/internal/model"
	// "go.uber.org/zap"
)

func (s *service) List(ctx context.Context) ([]*model.User, error) {
	// out, err := s.repo.CallExternalApi(ctx, "geolocate", in.Lat, in.Lng)
	// if err != nil {
	// 	s.logger.Error("search geocode failed", zap.Error(err))
	// 	return model.GeocodeOut{}, err
	// }

	var users []*model.User

	return users, nil
}
