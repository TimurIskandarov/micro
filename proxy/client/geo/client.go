package geo

import (
	"fmt"
	"proxy/config"

	proto "geo/pkg/geo_v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func New(conf config.RPCServer, logger *zap.Logger) proto.GeoV1Client {
	dsn := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	
	conn, err := grpc.Dial(dsn, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("grpc server connect error:", zap.Error(err))
	}

	return proto.NewGeoV1Client(conn)
}
