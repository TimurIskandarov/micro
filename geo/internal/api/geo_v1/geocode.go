package geo_v1

import (
	"context"
	"geo/internal/converter"

	proto "geo/pkg/geo_v1"
)

func (i *Implementation) GeoCode(ctx context.Context, req *proto.GeocodeRequest) (*proto.GeocodeResponse, error) {
	out, err := i.geoService.Geocode(ctx, *converter.ToGeocodeRequest(req))
	if err != nil {
		return nil, err
	}

	return converter.ToGeocodeResponseProto(out), nil
}
