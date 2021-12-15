package postgres

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mymhimself/ticket-system-api/internal/config"
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	"github.com/mymhimself/ticket-system-api/internal/repository"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgres struct {
	db     *gorm.DB
	logger logger.Logger
	cfg    config.DataProtocol
}

func New(cfg config.Config, logger logger.Logger) (repository.Postgres, error) {

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Postgres.Host, cfg.Postgres.Username, cfg.Postgres.DBName, cfg.Postgres.Password)
	conn, err := gorm.Open(driver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	//if err := conn.Ping(); err != nil {
	//	return nil, err
	//}

	return &postgres{db: conn, logger: logger, cfg: cfg.DataProtocol}, nil
}
