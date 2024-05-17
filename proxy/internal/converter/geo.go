package converter

import (
	"proxy/internal/model"

	proto "geo/pkg/geo_v1"
)

func ToSearchRequestProto(req *model.SearchIn) *proto.SearchRequest {
	return &proto.SearchRequest{
		Query: req.Query,
	}
}

func ToSearchOut(resp *proto.SearchResponse) *model.SearchOut {
	var addresses []*model.Address
	for _, address := range resp.GetAddresses() {
		addresses = append(addresses, &model.Address{
			Lat:   address.Lat,
			Lng:   address.Lon,
			Value: address.Value,
		})
	}

	return &model.SearchOut{
		Addresses: addresses,
		Status:    int(resp.GetStatus()),
	}
}

func ToGeocodeRequestProto(req *model.GeocodeIn) *proto.GeocodeRequest {
	return &proto.GeocodeRequest{
		Lat: req.Lat,
		Lon: req.Lng,
	}
}

func ToGeocodeOut(resp *proto.GeocodeResponse) *model.GeocodeOut {
	var addresses []*model.Address
	for _, address := range resp.GetAddresses() {
		addresses = append(addresses, &model.Address{
			Lat:   address.Lat,
			Lng:   address.Lon,
			Value: address.Value,
		})
	}

	return &model.GeocodeOut{
		Addresses: addresses,
		Status:    int(resp.GetStatus()),
	}
}
