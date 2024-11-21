package token

import "time"

type TokenMaker interface {
	CreateToken(username, roleId string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
