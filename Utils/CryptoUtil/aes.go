package CryptoUtil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

type CryptoAESUtils struct {
	key []byte
}

func NewCryptoAESUtils(key []byte) *CryptoAESUtils {
	return &CryptoAESUtils{key: key}
}

func (cu *CryptoAESUtils) EncryptJSON(jsonData []byte) ([]byte, error) {
	block, err := aes.NewCipher(cu.key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	encryptedData := aesGCM.Seal(nil, iv, jsonData, nil)

	result := make([]byte, len(iv)+len(encryptedData))
	copy(result, iv)
	copy(result[len(iv):], encryptedData)

	return result, nil
}

func (cu *CryptoAESUtils) DecryptJSON(encryptedData []byte) ([]byte, error) {
	if len(encryptedData) < aes.BlockSize {
		return nil, errors.New("invalid encrypted data")
	}

	iv := encryptedData[:aes.BlockSize]
	encryptedData = encryptedData[aes.BlockSize:]

	block, err := aes.NewCipher(cu.key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	decryptedData, err := aesGCM.Open(nil, iv, encryptedData, nil)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}
