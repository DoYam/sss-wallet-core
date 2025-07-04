package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/hashicorp/vault/shamir"
)

var Shares [][]byte
var PubKey *ecdsa.PublicKey

func InitKeyShares() error {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	PubKey = &priv.PublicKey
	shares, err := shamir.Split(priv.D.Bytes(), 5, 3)
	if err != nil {
		return err
	}

	Shares = shares
	return nil
}
