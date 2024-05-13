package geo_v1

import (
	"geo/internal/service/geo"
	proto "geo/pkg/geo_v1"
)

type Implementation struct {
	geoService geo.Service
	proto.UnimplementedGeoV1Server
}

func NewImplementation(geoService geo.Service) *Implementation {
	return &Implementation{
		geoService: geoService,
	}
}
