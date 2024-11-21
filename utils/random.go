package utils

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(lengthStringBuilder int) string {
	var substrings strings.Builder

	lengthAlphabet := len(alphabet)
	for i := 0; i < lengthStringBuilder; i++ {
		c := alphabet[rand.IntN(lengthAlphabet)]
		substrings.WriteByte(c)
	}

	return substrings.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
