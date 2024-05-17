package geo

import (
	"context"

	"proxy/internal/model"

	proto "geo/pkg/geo_v1"
)

var _ Service = (*service)(nil)

type Service interface {
	AddressSearch(ctx context.Context, in *model.SearchIn) (*model.SearchOut, error)
	Geocode(ctx context.Context, in *model.GeocodeIn) (*model.GeocodeOut, error)
}

type service struct {
	client proto.GeoV1Client
}

func NewService(client proto.GeoV1Client) *service {
	return &service{
		client: client,
	}
}
