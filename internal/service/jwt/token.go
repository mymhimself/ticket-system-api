package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/pkg/myerror"
	"time"
)

func (a jwtService) GenerateTokenPair(user *model.User) (string, string, error) {
	accessTokenExpireTime := time.Now().Add(time.Duration(a.cfg.AccessTokenExpireTime) * time.Minute)
	//initialize claim of access token
	claim := model.TokenClaims{
		ID:       user.Model.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpireTime.Unix(),
		},
	}

	//initialize claim of refresh token
	refreshTokenExpireTime := time.Now().Add(time.Duration(a.cfg.RefreshTokenExpireTime) * time.Minute)
	rtClaim := model.RefreshTokenClaims{
		ID: user.Model.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpireTime.Unix(),
		},
	}

	jwtAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedAccessToken, err := jwtAccessToken.SignedString([]byte(a.cfg.Secret))

	if err != nil {
		a.logger.Error(err.Error())
		return "", "", myerror.New(myerror.InternalError, enum.ServiceLayer, err.Error())
	} else {
		jwtRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaim)
		signedRefreshToken, err := jwtRefreshToken.SignedString([]byte(a.cfg.Secret))
		if err != nil {
			a.logger.Error(err.Error())
			return "", "", myerror.New(myerror.InternalError, enum.ServiceLayer, err.Error())
		}

		return signedAccessToken, signedRefreshToken, nil
	}
}
