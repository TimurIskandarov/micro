package config

import (
	"fmt"
	"os"
)

type AppConf struct {
	AppName   string    `yaml:"app_name"`
	Logger    Logger    `yaml:"logger"`
	RPCServer RPCServer `yaml:"rpc_server"`
	Token     Token     `yaml:"token"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type RPCServer struct {
	Port string `yaml:"port"`
	Type string `yaml:"type"`
}

type Token struct {
	AccessSecret string `yaml:"access_secret"`
}

func NewAppConf() AppConf {
	fmt.Println("ACCESS_SECRET", os.Getenv("ACCESS_SECRET"))
	fmt.Println("RPC_PORT", os.Getenv("RPC_PORT"))

	return AppConf{
		AppName: os.Getenv("APP_NAME"),
		Logger: Logger{
			Level: os.Getenv("LOG_LEVEL"),
		},
		RPCServer: RPCServer{
			Port: os.Getenv("RPC_PORT"),
			Type: os.Getenv("RPC_PROTOCOL"),
		},
		Token: Token{
			AccessSecret: os.Getenv("ACCESS_SECRET"),
		},
	}
}
