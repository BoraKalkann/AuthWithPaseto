package token

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoMaker struct {
	pasetoKey paseto.V4SymmetricKey
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != 32 {
		return nil, fmt.Errorf("geçersiz anahtar boyutu: 32 karakter olmalı, girilen: %d", len(symmetricKey))
	}

	key, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))
	if err != nil {
		return nil, err
	}

	return &PasetoMaker{
		pasetoKey: key,
	}, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	token := paseto.NewToken()

	now := time.Now()
	token.SetIssuedAt(now)
	token.SetNotBefore(now)
	token.SetExpiration(now.Add(duration))

	token.SetString("username", username)

	return token.V4Encrypt(maker.pasetoKey, nil), nil
}

func (maker *PasetoMaker) VerifyToken(tokenString string) (*paseto.Token, error) {
	parser := paseto.NewParser()

	token, err := parser.ParseV4Local(maker.pasetoKey, tokenString, nil)
	if err != nil {
		return nil, fmt.Errorf("geçersiz token: %w", err)
	}

	return token, nil
}
