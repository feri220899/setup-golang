package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func NormalizeKey(k string) string {
	if len(k) >= 16 {
		return k[:16]
	}
	return k + strings.Repeat("0", 16-len(k))
}

func Encrypt(plainText, key string) (string, error) {
	key = NormalizeKey(key)
	if len(key) != 16 {
		return "", errors.New("kunci harus tepat 16 karakter")
	}
	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, []byte(plainText))

	final := append(iv, cipherText...) // gabungkan IV + hasil enkripsi
	return base64.StdEncoding.EncodeToString(final), nil
}

func Decrypt(cipherTextB64, key string) (string, error) {
	key = NormalizeKey(key)
	if len(key) != 16 {
		return "", errors.New("kunci harus tepat 16 karakter")
	}
	keyBytes := []byte(key)

	data, err := base64.StdEncoding.DecodeString(cipherTextB64)
	if err != nil {
		return "", err
	}

	if len(data) < aes.BlockSize {
		return "", errors.New("data tidak valid (kurang dari 16 byte)")
	}

	iv := data[:aes.BlockSize]
	cipherText := data[aes.BlockSize:]

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}
