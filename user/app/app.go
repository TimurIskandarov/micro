package app

import (
	"context"
	"net"

	"user/config"
	"user/db"

	userV1 "user/internal/api/user_v1"
	userRepository "user/internal/repository/user"
	userService "user/internal/service/user"
	proto "user/pkg/user_v1"

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

	// Инициализация пула подключений к Postgres
	dbc := db.NewSqlDB(context.Background(), a.config.DB, a.logger)

	userRepo := userRepository.NewRepository(dbc)
	userSrv := userService.NewService(userRepo, a.logger)
	proto.RegisterUserV1Server(s, userV1.NewImplementation(userSrv))

	a.logger.Info("server started")

	if err := s.Serve(lis); err != nil {
		a.logger.Fatal("failed to serve:", zap.Error(err))
	}
}
