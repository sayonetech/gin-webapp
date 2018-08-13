package config

import (
	"go-webapp/common"
)

// Environment configuration file
// Multiple environment configurations can be configured for switching

const MaxAge int = 365 * 24 * 60 * 60

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
	SESSION_KEY       string
	VERSION           string
	REDIS_HOST        string
	SESSION_OBJ_KEY   string
	SENTRY_URL        string
	SET               bool // this flag is to show that envrionment struct has been set once,			      // the flag is set in code
}

var environs map[string]interface{}
var enviornment Env

func init() {
	if !enviornment.SET { // this needs to run only once
		environs = Read()
		enviornment = Env{
			DEBUG: common.Getenv("DEBUG"),

			SERVER_PORT:       string(environs["SERVER_PORT"].(string)),
			HOST:              string(environs["HOST"].(string)),
			DATABASE_PORT:     string(environs["DATABASE_PORT"].(string)),
			DATABASE_USERNAME: string(environs["DATABASE_USERNAME"].(string)),
			DATABASE_PASSWORD: string(environs["DATABASE_PASSWORD"].(string)),
			DATABASE_NAME:     string(environs["DATABASE_NAME"].(string)),

			ACCESS_LOG:      common.Getenv("ACCESS_LOG"),
			ACCESS_LOG_PATH: string(environs["ACCESS_LOG_PATH"].(string)),

			ERROR_LOG:      common.Getenv("ERROR_LOG"),
			ERROR_LOG_PATH: string(environs["ERROR_LOG_PATH"].(string)),

			APP_SECRET:      string(environs["APP_SECRET"].(string)),
			SESSION_KEY:     string(environs["SESSION_KEY"].(string)),
			VERSION:         string(environs["VERSION"].(string)),
			REDIS_HOST:      string(environs["REDIS_HOST"].(string)),
			SESSION_OBJ_KEY: string(environs["SESSION_OBJ_KEY"].(string)),
			SENTRY_URL:      string(environs["SENTRY_URL"].(string)),
			SET:             true,
		}
	}
	return
}

func GetEnvirons() *map[string]interface{} {
	return &environs
}

var sessionConfig = Config{
	Secret:   []byte(GetEnv().APP_SECRET),
	Name:     GetEnv().SESSION_KEY,
	Path:     "",
	Domain:   GetEnv().HOST,
	MaxAge:   MaxAge,
	Secure:   false,
	HttpOnly: false,
}

//GetEnv get the current enviornment configuration
func GetEnv() *Env {
	return &enviornment
}

//GetSessionConfig get the current session configuration
func GetSessionConfig() *Config {
	return &sessionConfig
}
