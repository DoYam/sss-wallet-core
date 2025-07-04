package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"

	"github.com/hashicorp/vault/shamir"
)

func RebuildPrivateKey(selected [][]byte) (*ecdsa.PrivateKey, error) {
	dBytes, err := shamir.Combine(selected)
	if err != nil {
		return nil, err
	}

	d := new(big.Int).SetBytes(dBytes)
	return &ecdsa.PrivateKey{
		D: d,
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     PubKey.X,
			Y:     PubKey.Y,
		},
	}, nil
}
