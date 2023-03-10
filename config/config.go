package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort uint
	DBUser     string
	DBPass     string
	DBHost     string
	DBPort     uint
	DBName     string
	JWTSecret  string
}

func NewConfig() *AppConfig {
	cfg := initConfig()
	if cfg == nil {
		log.Fatal("cannot run configuration setup")
		return nil
	}

	return cfg
}

func initConfig() *AppConfig {
	var app AppConfig

	godotenv.Load("config.env")

	serverPortConv, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("error parse server port")
		return nil
	}
	app.ServerPort = uint(serverPortConv)

	app.DBUser = os.Getenv("DB_USERNAME")

	app.DBPass = os.Getenv("DB_PASSWORD")

	app.DBHost = os.Getenv("DB_HOST")

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("error parse db port")
		return nil
	}
	app.DBPort = uint(port)

	app.DBName = os.Getenv("DB_NAME")

	app.JWTSecret = os.Getenv("JWT_SECRET")

	return &app
}
