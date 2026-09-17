package gateway

import (
	"crypto/sha256"
	"fmt"
)

type CryptoGw interface {
	GenerateHash(bytes []byte) string
}

type CryptoGateway struct{}

func NewCryptoGateway() *CryptoGateway {
	return &CryptoGateway{}
}

func (gw *CryptoGateway) GenerateHash(bytes []byte) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", bytes)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
