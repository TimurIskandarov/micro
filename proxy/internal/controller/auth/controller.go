package auth

import (
	"net/http"

	"proxy/internal/service/auth"

	"go.uber.org/zap"
)

var _ Controller = (*controller)(nil)

type Controller interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service auth.Service
	logger *zap.Logger
}

func NewController(service auth.Service, logger *zap.Logger) *controller {
	return &controller{
		service: service,
		logger: logger,
	}
}
