package postgres

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/pkg/myerror"
)

func (p *postgres) AutoMigration() error {
	err := p.db.AutoMigrate(
		&model.TicketThread{},
		&model.TicketMessage{},
		&model.User{},
	)
	if err != nil {
		return myerror.New(myerror.InternalError, enum.RepoLayer, err.Error())
	} else {
		p.fakePopulation()
		return nil
	}
}

func (p *postgres) fakePopulation() {
	//insert fake data, multiple admin and user
}
