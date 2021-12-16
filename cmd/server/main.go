package main

import (
	"fmt"
	"github.com/arkannsk/gql-example-service/config"
	"github.com/arkannsk/gql-example-service/db"
	"github.com/arkannsk/gql-example-service/server"
	auth_service "github.com/arkannsk/gql-example-service/services/auth-service"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logLevel, _ := log.ParseLevel(config.Param.LogLevel)
	log.SetLevel(logLevel)

	postgresDB := db.NewDBConnection(config.Param.DB)

	authService := auth_service.NewAuthService(postgresDB, config.Param.PhoneAuthCodeTTL,
		config.Param.PhoneAuthMaxWrongAttempts, config.Param.JWTTokenTTL,
		config.Param.JWTPublicKeyPath, config.Param.JWTPrivateKeyPath)

	sigChan := make(chan error, 2)
	go func() {
		log.WithFields(log.Fields{
			"transport": "http",
			"address":   ":" + config.Param.AppPort,
		}).Info("server listening")

		srv := server.NewServer(postgresDB, authService)
		sigChan <- srv.ListenAndServe(config.Param.AppPort)
		srv.Shutdown()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sigChan <- fmt.Errorf("received signal %s", <-c)
	}()

	log.Info("terminated ", <-sigChan)

}
