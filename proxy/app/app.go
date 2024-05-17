package app

import (
	"net/http"
	"time"

	"proxy/config"
	"proxy/public"
	"proxy/server"

	"proxy/client"

	"proxy/internal/controller/auth"
	"proxy/internal/controller/geo"

	authService "proxy/internal/service/auth"
	geoService "proxy/internal/service/geo"
	
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type App struct {
	conf   config.AppConf
	logger *zap.Logger
}

func NewApp(conf config.AppConf, logger *zap.Logger) *App {
	return &App{conf: conf, logger: logger}
}

func (a *App) Run() {
	client := client.New(a.conf, a.logger)

	authSrv := authService.NewService(client.Auth)
	authController := auth.NewController(authSrv, a.logger)	
	
	geoSrv := geoService.NewService(client.Geo)
	geoController := geo.NewController(geoSrv, a.logger)

	r := chi.NewRouter()

	r.Get("/swagger", public.SwaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/swagger.json")
	})

	r.Post("/api/register", authController.Register)
	r.Post("/api/login", authController.Login)

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
