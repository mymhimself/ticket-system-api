package cmd

import (
	"fmt"
	"github.com/mymhimself/ticket-system-api/internal/config"
	"github.com/mymhimself/ticket-system-api/internal/network/http/echo"
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger/zap"
	"github.com/mymhimself/ticket-system-api/internal/repository/postgres"
	"github.com/mymhimself/ticket-system-api/internal/service/account"
	"github.com/mymhimself/ticket-system-api/internal/service/jwt"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
)

func startServer(c *cli.Context) error {
	cfg := new(config.Config)
	config.ReadYAML("config.yaml", cfg)

	f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	loggerInstance := zap.New(f, zapcore.ErrorLevel)
	postgresInstance, err := postgres.New(*cfg, loggerInstance)
	if err != nil {
		return err
	}
	authenticationServiceInstance := jwt.New(cfg.Auth, loggerInstance)
	accountServiceInstance := account.New(cfg.Account, authenticationServiceInstance, postgresInstance, loggerInstance)
	httpServiceInstance := echo.New(cfg.Auth, loggerInstance, accountServiceInstance)

	go func() {
		if err := httpServiceInstance.Start(cfg.App.Address); err != nil {
			loggerInstance.Error(fmt.Sprintf("error happen while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	if err := httpServiceInstance.Shutdown(); err != nil {
		fmt.Println("\nREST server doesn't shutdown in 10 seconds")
	}

	return nil
}
