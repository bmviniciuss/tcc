package encrypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"os"
)

type Encrypter interface {
	Encrypt(content []byte) (encryptedContent []byte, err error)
}

type encryter struct {
	encryptionKey []byte
}

func NewEncrypter(encryptionKey []byte) *encryter {
	return &encryter{
		encryptionKey: encryptionKey,
	}
}

func (e *encryter) Encrypt(content []byte) (encryptedContent []byte, err error) {
	plainText := PKCS5Padding(content, aes.BlockSize)
	ciphertext := make([]byte, len(plainText))

	block, err := aes.NewCipher(e.encryptionKey)

	if err != nil {
		return nil, err
	}

	IV := []byte(os.Getenv("ENCRYPTION_IV"))
	mode := cipher.NewCBCEncrypter(block, IV)
	mode.CryptBlocks(ciphertext, plainText)
	encoded := b64.StdEncoding.EncodeToString(ciphertext)
	return []byte(encoded), nil
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
