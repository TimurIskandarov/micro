package app

import (
	"fmt"
	"net/http"
	"time"

	"proxy/config"
	"proxy/internal/controller/geo"
	geoService "proxy/internal/service/geo"
	proto "proxy/pkg/geo_v1"
	"proxy/server"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	conf   config.AppConf
	logger *zap.Logger
}

func NewApp(conf config.AppConf, logger *zap.Logger) *App {
	return &App{conf: conf, logger: logger}
}

func (a *App) Run() {
	dsn := fmt.Sprintf("%s:%s", a.conf.GeoServer.Host, a.conf.GeoServer.Port)

	conn, err := grpc.Dial(dsn, grpc.WithInsecure())
	if err != nil {
		a.logger.Fatal("grpc server connect error:", zap.Error(err))
	}

	client := proto.NewGeoV1Client(conn)
	geoSrv := geoService.NewService(client)
	geoController := geo.NewController(geoSrv, a.logger)

	r := chi.NewRouter()
	
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API"))
	})

	r.Group(func(r chi.Router) {
		r.Post("/api/address/search", geoController.AddressSearch)
		r.Post("/api/address/geocode", geoController.Geocode)
	})

	srv := &http.Server{
		Addr:         ":" + a.conf.Server.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := server.NewHttpServer(a.conf.Server, srv, a.logger)
	server.Serve()
}
