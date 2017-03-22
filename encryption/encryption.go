package encryption

import (
	"secureStore/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"log"
	"io"
)

var secret []byte = []byte(config.Parameters.Secret)
var nonce []byte  = make([]byte, 12)

func init() {
	// Initialize nonce value
	_, err := io.ReadFull(rand.Reader, nonce)
	checkError(err)
}

func Encrypt(value string) string {
	plaintext   := []byte(value)
	block, err  := aes.NewCipher(secret); checkError(err)
	aesgcm, err := cipher.NewGCM(block); checkError(err)
	ciphertext  := aesgcm.Seal(nil, nonce, plaintext, nil)
	finalOutput := base64.StdEncoding.EncodeToString(ciphertext)

	return finalOutput
}

func Decrypt(value string) string {
	// Decode Base64
	decoded, err := base64.StdEncoding.DecodeString(value); checkError(err)
	ciphertext := []byte(decoded)

	// Decrypt
	block, err      := aes.NewCipher(secret); checkError(err)
	aesgcm, err     := cipher.NewGCM(block); checkError(err)
	plaintext, err  := aesgcm.Open(nil, nonce, ciphertext, nil); checkError(err)

	return string(plaintext)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}