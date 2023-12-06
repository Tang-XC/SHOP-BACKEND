package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type userClaim struct {
	jwt.RegisteredClaims
}

var sercet = "SET YOUR SERCET HERE"

func GenerateToken(account string) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, &userClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 24),
			},
			Issuer:  "back-end",
			Subject: account,
		},
	})

	token, err := tokenStruct.SignedString([]byte(sercet))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*userClaim, error) {
	tokenStruct, err := jwt.ParseWithClaims(token, &userClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sercet), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := tokenStruct.Claims.(*userClaim)
	if !ok {
		return nil, err
	}
	return claim, nil
}
