package strings

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

func RsaEncode(publicKey, originalData string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", errors.New("base64 error: " + err.Error())
	}
	pubKey, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return "", errors.New("pubKey error: " + err.Error())
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))
	return base64.StdEncoding.EncodeToString(encryptedData), err
}

func RsaDecode(privateKey, encryptedData string) (string, error) {
	encryptedDecodeBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	key, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}
	prvKey, _ := x509.ParsePKCS8PrivateKey(key)
	originalData, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey.(*rsa.PrivateKey), encryptedDecodeBytes)
	return string(originalData), err
}
