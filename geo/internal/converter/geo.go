package converter

import (
	"geo/internal/model"
	proto "geo/pkg/geo_v1"
)

func ToSearchRequest(req *proto.SearchRequest) *model.SearchIn {
	return &model.SearchIn{
		Query: req.GetQuery(),
	}
}

func ToSearchResponseProto(resp model.SearchOut) *proto.SearchResponse {
	var addresses []*proto.Address
	for _, address := range resp.Addresses {
		addresses = append(addresses, &proto.Address{
			Lat: address.Lat,
			Lon: address.Lng,
			Value: address.Value,
		})
	}

	return &proto.SearchResponse{
		Addresses: addresses,
		Status: int64(resp.Status),
	}
}

func ToGeocodeRequest(req *proto.GeocodeRequest) *model.GeocodeIn {
	return &model.GeocodeIn{
		Lat: req.Lat,
		Lng: req.Lon,
	}
}

func ToGeocodeResponseProto(resp model.GeocodeOut) *proto.GeocodeResponse {
	var addresses []*proto.Address
	for _, address := range resp.Addresses {
		addresses = append(addresses, &proto.Address{
			Lat: address.Lat,
			Lon: address.Lng,
			Value: address.Value,
		})
	}
	
	return &proto.GeocodeResponse{
		Addresses: addresses,
		Status: int64(resp.Status),
	}
}