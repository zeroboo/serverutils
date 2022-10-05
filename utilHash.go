package serverutils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

type hashUtil struct {
}

func CreateSHA256(value string) string {
	hasher := sha256.New()
	hasher.Write([]byte(value))
	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func CreateMD5Hash(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateHMAC_SHA256(message string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(expectedMAC)
}

func CreateScrypt(password string, salt string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	CreateScrypt := base64.StdEncoding.EncodeToString(dk)
	return CreateScrypt, err
}
