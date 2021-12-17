package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Parameters struct {
	LogLevel                  string `envconfig:"LOG_LEVEL" default:"WARN"`
	AppPort                   string `envconfig:"APP_PORT" default:"8080"`
	ENV                       string `envconfig:"ENV" default:"production"`
	DB                        string `envconfig:"POSTGRES_URL" default:""`
	JWTPublicKeyPath          string `envconfig:"JWT_PUBLIC_KEY_PATH" default:""`
	JWTPrivateKeyPath         string `envconfig:"JWT_PRIVATE_KEY_PATH" default:""`
	JWTTokenTTL               int    `envconfig:"JWT_TOKEN_TTL" default:"120"`                 // in minutes
	PhoneAuthCodeTTL          int    `envconfig:"PHONE_AUTH_CODE_TTL" default:"120"`           // in seconds
	PhoneAuthMaxAttemptsCount int    `envconfig:"PHONE_AUTH_MAX_WRONG_ATTEMPTS" default:"120"` // in seconds
}

var Param Parameters

func init() {
	if err := godotenv.Load(); err != nil {
		log.Trace(err.Error())
	}
	if err := envconfig.Process("", &Param); err != nil {
		log.Trace(err.Error())
	}
}
