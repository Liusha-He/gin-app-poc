package auth

import "time"

type Maker interface {
	CreateToken(string, time.Duration) (string, error)
	VerifyToken(string) (*Payload, error)
}
