package app

import (
	"net"

	"user/config"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	config config.AppConf
	logger *zap.Logger
}

func NewApp(config config.AppConf, logger *zap.Logger) *App {
	return &App{
		config: config,
		logger: logger,
	}
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", ":"+a.config.RPCServer.Port)
	if err != nil {
		a.logger.Fatal("failed to listen: %s", zap.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	a.logger.Info("server started")

	if err := s.Serve(lis); err != nil {
		a.logger.Fatal("failed to serve:", zap.Error(err))
	}
}