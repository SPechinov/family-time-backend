package utils

import (
	"crypto/rand"
	mathrand "math/rand"
	"time"
)

func GenerateRandomCode(n int) string {
	const charset = "0123456789"
	code := make([]byte, n)
	randomBytes := make([]byte, n)

	if _, err := rand.Read(randomBytes); err != nil {
		// Use fallback
		rnd := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
		for i := range code {
			code[i] = charset[rnd.Intn(len(charset))]
		}
		return string(code)
	}

	for i, b := range randomBytes {
		code[i] = charset[b%uint8(len(charset))]
	}

	return string(code)
}
