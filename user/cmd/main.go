package main

import (
	"os"

	"user/app"
	"user/config"
	"user/logger"

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
