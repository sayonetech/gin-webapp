package config

import (
	"go-webapp/common"
	"os"
)

// Environment configuration file
// Multiple environment configurations can be configured for switching

//Env Enviornment config
type Env struct {
	DEBUG             bool
	HOST              string
	DATABASE_PORT     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
	SERVER_PORT       string
	ACCESS_LOG        bool
	APP_SECRET        string
	ACCESS_LOG_PATH   string
	ERROR_LOG         bool
	ERROR_LOG_PATH    string
}

var enviornment = Env{
	DEBUG: common.Getenv("DEBUG"),

	SERVER_PORT:       os.Getenv("SERVER_PORT"),
	HOST:              os.Getenv("HOST"),
	DATABASE_PORT:     os.Getenv("DATABASE_PORT"),
	DATABASE_USERNAME: os.Getenv("DATABASE_USERNAME"),
	DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
	DATABASE_NAME:     os.Getenv("DATABASE_NAME"),

	ACCESS_LOG:      common.Getenv("ACCESS_LOG"),
	ACCESS_LOG_PATH: os.Getenv("ACCESS_LOG_PATH"),

	ERROR_LOG:      common.Getenv("ERROR_LOG"),
	ERROR_LOG_PATH: os.Getenv("ERROR_LOG_PATH"),

	APP_SECRET: os.Getenv("APP_SECRET"),
}

//GetEnv get the current enviornment configuration
func GetEnv() *Env {
	return &enviornment
}
