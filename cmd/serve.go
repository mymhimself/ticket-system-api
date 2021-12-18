package cmd

import (
	"fmt"
	"github.com/mymhimself/ticket-system-api/internal/config"
	"github.com/mymhimself/ticket-system-api/internal/network/http/echo"
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger/zap"
	"github.com/mymhimself/ticket-system-api/internal/repository/postgres"
	"github.com/mymhimself/ticket-system-api/internal/service/account"
	"github.com/mymhimself/ticket-system-api/internal/service/jwt"
	"github.com/mymhimself/ticket-system-api/internal/service/ticket"
	"github.com/mymhimself/ticket-system-api/internal/service/validation"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
)

func startServer(c *cli.Context) error {
	//load config file
	cfg, cfgErr := config.ReadJSON("build/config/config.json")
	if cfgErr != nil {
		return cfgErr
	}
	//setting output log file
	logFile, logErr := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if logErr != nil {
		return logErr
	}

	//instantiate a logger service
	loggerInstance := zap.New(logFile, zapcore.ErrorLevel)
	databaseInstance, dbErr := postgres.New(cfg.Database.Postgres, cfg.DataProtocol, loggerInstance)
	if dbErr != nil {
		return dbErr
	}
	if cfg.Database.Postgres.AutoMigration {
		amErr := databaseInstance.AutoMigration()
		if amErr != nil {
			log.Fatal("Migration of database model failed.")
		}
	}

	//instantiate an authentication service
	authenticationServiceInstance := jwt.New(cfg.Authentication, loggerInstance)
	//instantiate an account service
	accountServiceInstance := account.New(authenticationServiceInstance, databaseInstance, loggerInstance)
	//instantiate a ticketing
	ticketServiceInstance := ticket.New(loggerInstance, databaseInstance)
	//instantiate a validation service
	validationServiceInstance := validation.New(loggerInstance)
	//instantiate a http server that implemented with echo framework
	httpServiceInstance := echo.New(cfg.Authentication, cfg.Authorization, loggerInstance, accountServiceInstance, ticketServiceInstance, validationServiceInstance)

	//run http server in new goroutine
	go func() {
		if err := httpServiceInstance.Start(cfg.App.Port); err != nil {
			loggerInstance.Error(fmt.Sprintf("error happen while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...\nThank you for using my Ticketing system.")

	if err := httpServiceInstance.Shutdown(); err != nil {
		fmt.Println("\nREST server doesn't shutdown in 10 seconds")
	}

	return nil
}
