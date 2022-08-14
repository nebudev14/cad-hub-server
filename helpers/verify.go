package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Decrypt(data []byte, tag []byte) []byte {
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println(err)
	}

	aesKey := []byte(os.Getenv("AES_KEY"))
	nonce := []byte(os.Getenv("IV"))
	aad := []byte(os.Getenv("AAD"))

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		fmt.Println(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}

	plaintext, err := aesgcm.Open(tag, nonce, data, aad)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}