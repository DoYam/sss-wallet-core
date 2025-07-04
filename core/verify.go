package core

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"
)

func Verify(msg string, r, s *big.Int, pub *ecdsa.PublicKey) bool {
	hash := sha256.Sum256([]byte(msg))
	return ecdsa.Verify(pub, hash[:], r, s)
}
