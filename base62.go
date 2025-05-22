package main

import (
	"crypto/sha256"
	"math/big"
	"time"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func base62Encode(data []byte) string {
	num := new(big.Int).SetBytes(data)
	var result string
	base := int64(len(charset))

	for num.Sign() > 0 {
		rem := new(big.Int)
		num.DivMod(num, big.NewInt(base), rem)
		result = string(charset[rem.Int64()]) + result
	}
	return result
}

func encodeShortCode(url string, id int64) string {
	timestamp := time.Now().UnixNano()

	// Combine URL + timestamp + ID
	input := url + string(timestamp) + string(id)

	// Hash the input
	hash := sha256.Sum256([]byte(input))

	// Encode hash to base62
	base62 := base62Encode(hash[:])

	// Return only first 7 characters
	if len(base62) > 7 {
		return base62[:7]
	}
	return base62
}
