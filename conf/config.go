package conf

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

var Cfg AppCfg

func ReadEnv() {
	err := godotenv.Load()
	if err != nil {
		//l.Error.Println("error reading .env file")
	}

	Cfg.AppAddr = os.Getenv("APP_ADDR")
	Cfg.AppPort = os.Getenv("APP_PORT")

	Cfg.DbHost = os.Getenv("POSTGRES_HOST")
	Cfg.DbPort = os.Getenv("POSTGRES_PORT")
	Cfg.DbName = os.Getenv("POSTGRES_DB")
	Cfg.DbUser = os.Getenv("POSTGRES_USER")
	Cfg.DbPassword = os.Getenv("POSTGRES_PASSWORD")

}
