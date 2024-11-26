package token

import (
	"encoding/json"
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoMaker struct {
	privateKey paseto.V4AsymmetricSecretKey
	parser     paseto.Parser
}

func NewPasetoMaker() TokenMaker {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	parser := paseto.NewParserWithoutExpiryCheck()

	return &PasetoMaker{
		privateKey: privateKey,
		parser:     parser,
	}
}

func (maker *PasetoMaker) CreateToken(username string, roleId int, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, roleId, duration)
	if err != nil {
		return "", err
	}
	claims, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	token, err := paseto.NewTokenFromClaimsJSON(claims, nil)
	if err != nil {
		return "", err
	}

	tokenSigned := token.V4Sign(maker.privateKey, nil)
	return tokenSigned, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	tokenParser, err := maker.parser.ParseV4Public(maker.privateKey.Public(), token, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(tokenParser.ClaimsJSON(), payload)
	if err != nil {
		return nil, err
	}

	if err = payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
