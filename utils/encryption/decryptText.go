package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

const DECRYPTION_KEY = "adminadminadminadminadminadmin12"

func Decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(ENCRYPTION_KEY))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
