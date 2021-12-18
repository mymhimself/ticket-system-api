package postgres

import (
	"errors"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/pkg/myerror"
	"gorm.io/gorm"
)

func (p *postgres) CreateUser(user *model.User) error {
	//searching user data in database
	result := p.db.Create(&user)
	if result.Error != nil {
		//database error
		return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
	} else { //user data found
		return nil
	}
}

func (p *postgres) UpdateUser(user *model.User) error {
	result := p.db.Save(user)
	if result.Error != nil {
		return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
	} else {
		return nil
	}
}

func (p *postgres) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	//searching user data in database
	result := p.db.Where(id).First(&user)
	if result.Error != nil {
		//user data not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
			//database error
		} else {
			return nil, myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
		}
	} else { //user data found
		return &user, nil
	}
}

func (p *postgres) GetUserByUsername(username string) (*model.User, bool, error) {
	var user model.User
	//searching user data in database
	result := p.db.Where(model.User{Username: username}).First(&user)
	if result.Error != nil {
		//user data not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, false, nil
			//database error
		} else {
			err := myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
			p.logger.Error(err.Code.String(), err)
			return nil, false, err
		}
	} else { //user data found
		return &user, true, nil
	}
}

func (p *postgres) DeleteUser(id uint) error {
	result := p.db.Where(id).Delete(&model.User{})
	if result.Error != nil {
		return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
	} else {
		return nil
	}
}

func (p *postgres) CheckUserExistence(username string) (bool, error) {
	result := p.db.Where(model.User{Username: username}).First(&model.User{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
		}
	} else {
		return true, nil
	}
}
