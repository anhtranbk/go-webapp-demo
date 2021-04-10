package stringutil

import "math/rand"

const (
	AsciiUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AsciiLowerCase = "abcdefghijklmnopqrstuvwxyz"
	AsciiLetters   = AsciiLowerCase + AsciiUpperCase
	Digits         = "1234567890"
)

var letters = []rune(AsciiLetters + Digits)

func RandStringSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
