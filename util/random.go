package util

import "crypto/rand"

const PotentialBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyx_"

func RandomString(n int) string {

	if n <= 0 {
		return ""
	}

	var bytes = make([]byte, n)

	rand.Read(bytes)

	for i, b := range bytes {
		bytes[i] = PotentialBytes[b%byte(len(PotentialBytes))]
	}

	return string(bytes)
}
