package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/pkg/myerror"
	"time"
)

func (a jwtService) GenerateToken(user *model.User) (string, error) {
	expireTime := time.Now().Add(time.Duration(a.cfg.TokenExpireTime) * time.Millisecond * 60)
	claim := model.TokenClaims{
		ID:       user.Model.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := jwtToken.SignedString([]byte(a.cfg.Secret))
	if err != nil {
		a.logger.Error(err.Error())
		return "", myerror.New(myerror.InternalError, enum.ServiceLayer, err.Error())
	} else {
		return signedToken, nil
	}
}

func (a jwtService) GenerateRefreshToken(user *model.User) (string, error) {
	expireTime := time.Now().Add(time.Duration(a.cfg.RefreshExpireTime) * time.Millisecond * 60)
	claim := model.TokenClaims{
		ID:       user.Model.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	signedToken, err := jwtToken.SignedString(a.cfg.Secret)
	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}
