package main

import (
	"os"

	"proxy/app"
	"proxy/config"
	"proxy/logger"

	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()

	// Создание конфигурации приложения
	conf := config.NewAppConf()
	// Создание логгера
	logger := logger.NewLogger(conf, os.Stdout)

	if err != nil {
		logger.Fatal("error loading .env file")
	}
	
	app := app.NewApp(conf, logger)
	app.Run()
}
