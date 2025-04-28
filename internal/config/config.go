package config

import (
	"github.com/joho/godotenv"
	"os"
)

type AppCfg struct {
	AppAddr string
	AppPort string

	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func ReadEnv() *AppCfg {
	var cfg AppCfg
	err := godotenv.Load()
	if err != nil {
		//l.Error.Println("error reading .env file")
	}

	cfg.AppAddr = os.Getenv("APP_ADDR")
	cfg.AppPort = os.Getenv("APP_PORT")

	cfg.DbHost = os.Getenv("POSTGRES_HOST")
	cfg.DbPort = os.Getenv("POSTGRES_PORT")
	cfg.DbName = os.Getenv("POSTGRES_DB")
	cfg.DbUser = os.Getenv("POSTGRES_USER")
	cfg.DbPassword = os.Getenv("POSTGRES_PASSWORD")
	return &cfg
}
