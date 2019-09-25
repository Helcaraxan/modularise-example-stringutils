package random

import (
	"math/rand"
	"strings"
)

func GenerateRandomNumber() int {
	return rand.Int()
}

func GenerateRandomString(length int) string {
	res := &strings.Builder{}
	for i := 0; i < length; i++ {
		res.WriteByte(byte('a' + rand.Int()%26))
	}
	return res.String()
}
