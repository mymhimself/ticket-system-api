package postgres

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mymhimself/ticket-system-api/internal/config"
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	"github.com/mymhimself/ticket-system-api/internal/repository"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type postgres struct {
	db              *gorm.DB
	logger          logger.Logger
	dataProtocolCfg config.DataProtocol
}

func New(cfg config.Postgres, dataProtocolCfg config.DataProtocol, logger logger.Logger) (repository.Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Host, cfg.Username, cfg.DBName, cfg.Password)
	conn, err := gorm.Open(driver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(DatabaseSetupFailed, err.Error())
	}

	return &postgres{db: conn, logger: logger, dataProtocolCfg: dataProtocolCfg}, nil
}
