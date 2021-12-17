package account

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/pkg/myerror"
)

func (a *accountService) Login(username, password string) (string, string, error) {
	// start span with context

	user, err := a.postgres.GetUserByUsername(username)
	if err != nil {
		return "", "", err
	}

	match := comparePassword(password, user.HashedPassword, user.Salt)
	if !match {
		return "", "", myerror.New(myerror.PasswordIncorrect, enum.ServiceLayer, "")
	} else {
		token, refreshToken, err := a.Authentication.GenerateTokenPair(user)
		return token, refreshToken, err
	}
}

func (a *accountService) Register(username, password string) (string, string, error) {
	var user model.User
	exists, err := a.postgres.CheckUserExistence(username)
	if err != nil {
		return "", "", err
	}
	if exists {
		return "", "", myerror.New(myerror.UserExists, enum.ServiceLayer, "")
	}

	user.Username = username
	user.Salt, err = generateRandomString(32)
	if err != nil {
		return "", "", myerror.New(myerror.InternalError, enum.ServiceLayer, err.Error())
	}

	user.HashedPassword = createHashedPassword(password, user.Salt)
	user.Role = enum.User
	err = a.postgres.CreateUser(&user)
	if err != nil {
		return "", "", err
	}
	accessToken, refreshToken, err := a.Authentication.GenerateTokenPair(&user)
	if err != nil {
		a.postgres.DeleteUser(user.ID)
		return "", "", myerror.New(myerror.InternalError, enum.ServiceLayer, err.Error())
	}

	return accessToken, refreshToken, nil
}

func (a *accountService) RegisterConfirm(key, phone, otp string) (string, error) {
	return "", nil
}
