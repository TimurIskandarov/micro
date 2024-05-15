package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type AppConf struct {
	AppName   string    `yaml:"app_name"`
	Logger    Logger    `yaml:"logger"`
	Server    Server    `yaml:"server"`
	GeoServer GeoServer `yaml:"geo_server"`
}

type Server struct {
	Port            string        `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type GeoServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}


type Logger struct {
	Level string `yaml:"level"`
}

func NewAppConf() AppConf {
	fmt.Println("SERVER_PORT", os.Getenv("SERVER_PORT"))
	fmt.Println("SHUTDOWN_TIMEOUT", os.Getenv("SHUTDOWN_TIMEOUT"))	
	fmt.Println("GEO_GRPC_HOST", os.Getenv("GEO_GRPC_HOST"))
	fmt.Println("GEO_GRPC_PORT", os.Getenv("GEO_GRPC_PORT"))

	return AppConf{
		AppName: os.Getenv("APP_NAME"),
		Logger: Logger{
			Level: os.Getenv("LOG_LEVEL"),
		},
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
		},		
		GeoServer: GeoServer{
			Host: os.Getenv("GEO_GRPC_HOST"),
			Port: os.Getenv("GEO_GRPC_PORT"),
		},
	}
}

func (a *AppConf) Init(logger *zap.Logger) {
	shutDownTimeOut, err := strconv.Atoi(os.Getenv("SHUTDOWN_TIMEOUT"))
	if err != nil {
		logger.Fatal("config: parse server shutdown timeout error")
	}
	shutDownTimeout := time.Duration(shutDownTimeOut) * time.Second
	a.Server.ShutdownTimeout = shutDownTimeout
}
