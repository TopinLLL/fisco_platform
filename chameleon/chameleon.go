package chameleon

import (
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

func Seal(privateKey string, data string) string {
	// size必须为32
	key := readSecretKey(chacha20poly1305.KeySize)
	c, _ := chacha20poly1305.New(key)
	nonce := readRandomNonce(c.NonceSize())

	return fmt.Sprintf("%x", c.Seal(nil, nonce, []byte(privateKey), []byte(data)))
}

func readSecretKey(i int) []byte {
	return make([]byte, i)
}

func readRandomNonce(i int) []byte {
	return make([]byte, i)
}
