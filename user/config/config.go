package config

import (
	"fmt"
	"os"
)

type AppConf struct {
	AppName   string    `yaml:"app_name"`
	Logger    Logger    `yaml:"logger"`
	DB        DB        `yaml:"db"`
	RPCServer RPCServer `yaml:"rpc_server"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type DB struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `json:"-" yaml:"user"`
	Password string `json:"-" yaml:"password"`
}

type RPCServer struct {
	Port string `yaml:"port"`
	Type string `yaml:"type"`
}

func NewAppConf() AppConf {
	fmt.Println("DB_HOST", os.Getenv("PG_HOST"))
	fmt.Println("DB_PORT", os.Getenv("PG_PORT"))
	fmt.Println("DB_NAME", os.Getenv("PG_DATABASE"))
	fmt.Println("DB_USER", os.Getenv("PG_USER"))
	fmt.Println("DB_PASSWORD", os.Getenv("PG_PASSWORD"))

	return AppConf{
		AppName: os.Getenv("APP_NAME"),
		Logger: Logger{
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: DB{
			Host:     os.Getenv("PG_HOST"),
			Port:     os.Getenv("PG_PORT"),
			Name:     os.Getenv("PG_DATABASE"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
		},
		RPCServer: RPCServer{
			Port: os.Getenv("RPC_PORT"),
			Type: os.Getenv("RPC_PROTOCOL"),
		},
	}
}
