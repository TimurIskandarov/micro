package geo

import (
	"net/http"

	"proxy/internal/service/geo"

	"go.uber.org/zap"
)

type Controller interface {
	AddressSearch(w http.ResponseWriter, r *http.Request)
	Geocode(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service geo.Service
	logger *zap.Logger
}

func NewController(service geo.Service, logger *zap.Logger) *controller {
	return &controller{
		service: service,
		logger: logger,
	}
}
