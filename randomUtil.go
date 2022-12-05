package serverutils

import (
	"math/rand"
	"time"
)

const CapitalLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NormalLetters string = "abcdefghijklmnopqrstuvwxyz"
const AllLetters string = NormalLetters + CapitalLetters
const Numbers string = "0123456789"

var seed int64 = time.Now().UnixNano()

type randomUtil struct {
}

func init() {
	rand.Seed(seed)
}

// GetRandomStringFromString generate random string
func (*randomUtil) GetRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = AllLetters[rand.Intn(len(AllLetters))]
	}
	return string(b)
}

// GetRandomStringFromString generate random string from given letters
func (*randomUtil) GetRandomStringFromString(n int, letters string) string {

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GetRandomStringUppercase generates random string with all upper case letters
func (*randomUtil) GetRandomStringUppercase(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = CapitalLetters[rand.Intn(len(CapitalLetters))]
	}
	return string(b)
}

// GetRandomNumbers return a random string of numbers
func (*randomUtil) GetRandomNumbers(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Numbers[rand.Intn(len(Numbers))]
	}
	return string(b)
}
