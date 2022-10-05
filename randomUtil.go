package serverutils

import (
	"math/rand"
	"time"
)

const capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const normalLetters = "abcdefghijklmnopqrstuvwxyz"
const allLetters = normalLetters + capitalLetters

var seed int64 = time.Now().UnixNano()

type randomUtil struct {
}

func init() {
	rand.Seed(seed)
}

func (*randomUtil) GetRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = allLetters[rand.Intn(len(allLetters))]
	}
	return string(b)
}

func (*randomUtil) GetRandomStringFromString(n int, letters string) string {

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (*randomUtil) GetRandomStringCapitalize(n int) string {

	b := make([]byte, n)
	for i := range b {
		b[i] = capitalLetters[rand.Intn(len(capitalLetters))]
	}
	return string(b)
}
