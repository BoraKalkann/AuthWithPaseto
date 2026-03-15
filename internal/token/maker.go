package token

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(tokenString string) (*paseto.Token, error)
}
