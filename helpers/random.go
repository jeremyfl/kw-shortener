package helpers

import "math/rand"

// GenerateRandomCode generate wth math
func GenerateRandomCode() string {
	len := 5

	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}

	return string(bytes)
}
