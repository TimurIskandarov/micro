package client

import (
	"proxy/client/auth"
	"proxy/client/geo"
	"proxy/config"

	"auth/pkg/auth_v1"
	"geo/pkg/geo_v1"

	"go.uber.org/zap"
)

type Client struct {
	Auth auth_v1.AuthV1Client
	Geo  geo_v1.GeoV1Client
}

func New(conf config.AppConf, logger *zap.Logger) *Client {
	return &Client{
		Auth: auth.New(conf.AuthServer, logger),
		Geo:  geo.New(conf.GeoServer, logger),
	}
}
