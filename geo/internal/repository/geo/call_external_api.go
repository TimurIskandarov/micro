package geo

import (
	"context"
	"fmt"
	"geo/internal/model"

	"github.com/ekomobile/dadata/v2/api/suggest"
)

func (r *repository) CallExternalApi(ctx context.Context, endpoint string, params ...string) ([]*model.Address, error) {
	var suggestions []*suggest.AddressSuggestion
	var err error

	switch endpoint {
	case "address":
		suggestions, err = r.callExternalApiByAddress(ctx, params...)
	case "geolocate":
		suggestions, err = r.callExternalApiByGeocode(ctx, params...)
	default:
		fmt.Println("endpoint not found")
	}

	if err != nil {
		return nil, err
	}

	addresses := make([]*model.Address, len(suggestions))
	for i, s := range suggestions {
		addresses[i] = &model.Address{
			Value: s.Value,
			Lat:   s.Data.GeoLat,
			Lng:   s.Data.GeoLon,
		}
	}

	return addresses, nil
}

func (r *repository) callExternalApiByAddress(ctx context.Context, params ...string) ([]*suggest.AddressSuggestion, error) {
	reqParams := suggest.RequestParams{Query: params[0]}
	suggestions, err := r.Api.Address(ctx, &reqParams)
	if err != nil {
		return nil, err
	}

	return suggestions, nil
}

func (r *repository) callExternalApiByGeocode(ctx context.Context, params ...string) ([]*suggest.AddressSuggestion, error) {
	// запрос
	reqParams := struct {Lat string `json:"lat"`; Lng string `json:"lon"`}{
		Lat: params[0],
		Lng: params[1],
	}
	// ответ
	result := struct {
		Suggestions []*suggest.AddressSuggestion
	}{
		Suggestions: nil,
	}

	err := r.Api.Client.Post(ctx, "geolocate/address", &reqParams, &result)
	if err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}
