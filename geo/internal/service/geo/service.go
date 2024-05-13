package geo

import (
	"context"

	"geo/internal/model"
	"geo/internal/repository/geo"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	AddressSearch(ctx context.Context, in model.SearchIn) (model.SearchOut, error)
	Geocode(ctx context.Context, in model.GeocodeIn) (model.GeocodeOut, error)
}

type service struct {
	repo   geo.Repository
	cache  *redis.Pool
	logger *zap.Logger
}

func NewService(repo geo.Repository, cache *redis.Pool, logger *zap.Logger) *service {
	return &service{
		repo:   repo,
		cache:  cache,
		logger: logger,
	}
}
