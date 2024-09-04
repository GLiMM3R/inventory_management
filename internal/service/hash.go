package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

func EncodeToSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func DecodeFromSHA256(input string) ([]byte, error) {
	return hex.DecodeString(input)
}

func EncodeToAES256(input string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(input)
	cfb := cipher.NewCFBEncrypter(block, make([]byte, aes.BlockSize))
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return hex.EncodeToString(ciphertext), nil
}

func DecodeFromAES256(input string, key []byte) (string, error) {
	ciphertext, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, make([]byte, aes.BlockSize))
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}
