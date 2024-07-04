package token

import "time"

type Maker interface {
	CreateToken(id int32, role string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
