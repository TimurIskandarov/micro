package app

import (
	"net"

	"auth/config"

	authV1 "auth/internal/api/auth_v1"
	authService "auth/internal/service/auth"
	proto "auth/pkg/auth_v1"

	"github.com/go-chi/jwtauth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	token  *jwtauth.JWTAuth
	config config.AppConf
	logger *zap.Logger
}

func NewApp(config config.AppConf, logger *zap.Logger) *App {
	return &App{
		token: jwtauth.New("HS256", []byte(config.Token.AccessSecret), nil),
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

	authSrv := authService.NewService(a.token, a.logger)
	proto.RegisterAuthV1Server(s, authV1.NewImplementation(authSrv))

	a.logger.Info("server started")

	if err := s.Serve(lis); err != nil {
		a.logger.Fatal("failed to serve:", zap.Error(err))
	}
}
