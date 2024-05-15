package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"geo/config"
	"geo/db"

	geoV1 "geo/internal/api/geo_v1"
	geoRepository "geo/internal/repository/geo"
	geoService "geo/internal/service/geo"
	proto "geo/pkg/geo_v1"

	"github.com/gomodule/redigo/redis"
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

	// Инициализация пула подключений к Redis
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			redisConn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				fmt.Println("redisConn не получен")
				return nil, err
			}
			return redisConn, err
		},
		TestOnBorrow: func(redisConn redis.Conn, _ time.Time) error {
			_, err := redisConn.Do("PING")
			fmt.Println("PING!")
			return err
		},
	}


	geoRepo := geoRepository.NewRepository(dbc)
	geoSrv := geoService.NewService(geoRepo, redisPool, a.logger)
	proto.RegisterGeoV1Server(s, geoV1.NewImplementation(geoSrv))

	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %s", err.Error())
	}
}
