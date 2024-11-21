package token

import (
	"fmt"
	"time"

	"github.com/go-http-server/core/utils"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTToken(secretKey string) (TokenMaker, error) {
	if len(secretKey) < utils.MIN_SIZE_SECRET_KEY {
		return nil, fmt.Errorf("Secret key to short, I prefer at least %d character.", utils.MIN_SIZE_SECRET_KEY)
	}

	return &JWTMaker{
		secretKey: secretKey,
	}
}

func (maker *JWTMaker) CreateToken(username, roleId string, duration time.Duration) (string, error) {
}
