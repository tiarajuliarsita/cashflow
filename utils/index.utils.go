package utils

import (
	"math/rand"
	"time"
)

var charset =[]byte("1234567890")
const prefix = "12345"
func RandomNumberAccount(n int) string {
	rand.Seed(time.Now().UnixMilli())
	number := make([]byte, n)
	for i := range number {
		number[i] = charset[rand.Intn(len(charset))]
	}

	return prefix + string(number)
}
