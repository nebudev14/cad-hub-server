package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"hex"
	"os"

	"github.com/joho/godotenv"
)

func Decrypt(data []byte) {
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println(err)
	}

	aesKey := []byte(os.Getenv("AES_KEY"))
	nonce := []byte(os.Getenv("IV"))
	aad := []byte(os.Getenv("AAD"))
}