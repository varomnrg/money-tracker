package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomID(length int) string {
	rand.Seed(time.Now().UnixNano())

	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	id := make([]byte, length)

	for i := 0; i < length; i++ {
		id[i] = charset[rand.Intn(len(charset))]
	}

	return string(id)
}
