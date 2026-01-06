package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfigStruct struct {
	AppPort           string
	AccessTokenSecret string
}

var AppConfig AppConfigStruct

func LoadEnv() {
	_ = godotenv.Load()

	AppConfig = AppConfigStruct{
		AppPort:           os.Getenv("APP_PORT"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}

	if AppConfig.AppPort == "" {
		log.Fatal("APP_PORT is required")
	}

	if AppConfig.AccessTokenSecret == "" {
		log.Fatal("ACCESS_TOKEN_SECRET is required")
	}
}
