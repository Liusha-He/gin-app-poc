package auth

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PASETOMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPASETOMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size, must be %d characters", chacha20poly1305.KeySize)
	}

	maker := PASETOMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (pm PASETOMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return pm.paseto.Encrypt(pm.symmetricKey, payload, nil)
}

func (pm PASETOMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := pm.paseto.Decrypt(token, pm.symmetricKey, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
