package token

import "time"

type Payload struct {
	Username  string    `json:"username"`
	RoleId    string    `json:"role_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
