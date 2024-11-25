package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-http-server/core/utils"
	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (TokenMaker, error) {
	if len(secretKey) < utils.MIN_SIZE_SECRET_KEY {
		return nil, fmt.Errorf("Secret key to short, I prefer at least %d character.", utils.MIN_SIZE_SECRET_KEY)
	}

	maker := &JWTMaker{
		secretKey: secretKey,
	}
	return maker, nil
}

func (maker *JWTMaker) CreateToken(username string, roleId int, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, roleId, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(utils.UNEXPECTED_SIGNING_METHOD_TOKEN + token.Method.Alg())
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New(utils.ERROR_CONVERT_TOKEN)
	}

	if err = payload.Valid(); err != nil {
		return nil, errors.New(utils.TOKEN_EXPIRED)
	}

	return payload, nil
}
