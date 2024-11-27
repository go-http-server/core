package utils

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	numeric  = "0123456789"
)

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

func RandomInt(from, to int) int {
	return from + rand.IntN(to-from+1)
}

func RandomCode() string {
	var substrings strings.Builder
	lengthNumeric := len(numeric)

	for i := 0; i < 6; i++ {
		c := numeric[rand.IntN(lengthNumeric)]
		substrings.WriteByte(c)
	}

	return substrings.String()
}
