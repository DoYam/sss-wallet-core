package core

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
)

type Signature struct {
	R string `json:"r"`
	S string `json:"s"`
}

func SignWithKey(priv *ecdsa.PrivateKey, msg string) (*Signature, error) {
	hash := sha256.Sum256([]byte(msg))
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return nil, err
	}
	return &Signature{
		R: r.Text(16),
		S: s.Text(16),
	}, nil
}
