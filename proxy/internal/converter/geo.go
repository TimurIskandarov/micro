package converter

import (
	"proxy/internal/model"
	proto "proxy/pkg/geo_v1"
)

func ToSearchRequestProto(req model.SearchIn) *proto.SearchRequest {
	return &proto.SearchRequest{
		Query: req.Query,
	}
}

func ToGeocodeRequestProto(req model.GeocodeIn) *proto.GeocodeRequest {
	return &proto.GeocodeRequest{
		Lat: req.Lat, 
		Lon: req.Lng,
	}
}
