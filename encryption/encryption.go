package encryption

import (
	"secureStore/config"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"io"
	"strings"
)

var secret []byte = []byte(config.Parameters.Secret)
var nonce []byte  = make([]byte, 12)

func Encrypt(value string, nonceValue string) string {
	// Encrypt
	_, err      := io.ReadFull(strings.NewReader(string(nonceValue)), nonce); checkError(err)
	plaintext   := []byte(value)
	block, err  := aes.NewCipher(secret); checkError(err)
	aesgcm, err := cipher.NewGCM(block); checkError(err)
	ciphertext  := aesgcm.Seal(nil, nonce, plaintext, nil)

	// Encode Base64
	finalOutput := base64.StdEncoding.EncodeToString(ciphertext)

	return finalOutput
}

func Decrypt(value string, nonceValue string) string {
	// Decode Base64
	decoded, err := base64.StdEncoding.DecodeString(value); checkError(err)
	ciphertext   := []byte(decoded)

	// Decrypt
	_, err          = io.ReadFull(strings.NewReader(string(nonceValue)), nonce); checkError(err)
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