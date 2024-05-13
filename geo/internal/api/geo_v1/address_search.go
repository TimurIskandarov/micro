package geo_v1

import (
	"context"

	"geo/internal/converter"
	proto "geo/pkg/geo_v1"
)

func (i *Implementation) AddressSearch(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	out, err := i.geoService.AddressSearch(ctx, *converter.ToSearchRequest(req))
	if err != nil {
		return nil, err
	}

	return converter.ToSearchResponseProto(out), nil
}
