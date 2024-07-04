package operator

import (
	"crypto/rand"
)

// generateRandomBytes returns a random [32]byte array.
func generateRandomBytes() ([32]byte, error) {
	var salt [32]byte
	_, err := rand.Read(salt[:]) // Read 32 bytes into the array
	if err != nil {
		return salt, err
	}
	return salt, nil
}
