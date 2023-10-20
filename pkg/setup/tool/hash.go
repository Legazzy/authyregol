package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

func NewHash(string string) string {
	byt := []byte(string)
	alg := sha256.New()

	alg.Write(byt)

	return hex.EncodeToString(alg.Sum(nil))
}
