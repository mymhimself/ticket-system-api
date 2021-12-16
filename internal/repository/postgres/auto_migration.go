package postgres

import "github.com/mymhimself/ticket-system-api/internal/entity/model"

func (p *postgres) AutoMigration() error {
	err := p.db.AutoMigrate(
		&model.Ticket{},
		&model.User{},
	)
	if err != nil {
		return err
	} else {
		return nil
	}
}
