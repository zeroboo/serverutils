package serverutils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

type hashUtil struct {
}

func (*hashUtil) CreateSHA256HashHex(data []byte) string {
	hasher := sha256.New()
	hasher.Write(data)
	sha256 := hex.EncodeToString(hasher.Sum(nil))
	return sha256
}

func (*hashUtil) CreateSHA256HashBase64(data []byte) string {
	hasher := sha256.New()
	hasher.Write(data)
	sha256 := base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
	return sha256
}

func (*hashUtil) CreateSHA1HashHex(data []byte) string {
	hasher := sha1.New()
	hasher.Write(data)
	//sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	sha := hex.EncodeToString(hasher.Sum(nil))

	return sha
}

func (*hashUtil) CreateSHA1HashBase64(data []byte) string {
	hasher := sha1.New()
	hasher.Write(data)
	//sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	sha := base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
func (*hashUtil) CreateMD5Hash(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (*hashUtil) CreateHMAC_SHA256(message string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(expectedMAC)
}

func (*hashUtil) CreateScrypt(password string, salt string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	CreateScrypt := base64.StdEncoding.EncodeToString(dk)
	return CreateScrypt, err
}
