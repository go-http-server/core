package token

import "time"

type TokenMaker interface {
	CreateToken(username string, roleId int, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
