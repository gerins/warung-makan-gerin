package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const ENCRYPTION_KEY = "adminadminadminadminadminadmin12"

func Encrypt(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(ENCRYPTION_KEY))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
