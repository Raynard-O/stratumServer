package main

import (
	"encoding/hex"
	"math/rand"
)


//randomHex returns a 64-bit random hex string.
func randomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
