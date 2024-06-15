package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// CalculateSHA256Checksum calculates the SHA256 checksum of a file
func CalculateSHA256Checksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := hash.Sum(nil)

	return hex.EncodeToString(checksum), nil
}

// CalculateSHA256Checksum calculates the SHA256 checksum of a []byte
func CalculateSHA256ChecksumBytes(data []byte) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write(data); err != nil {
		return "", err
	}

	checksum := hash.Sum(nil)

	return hex.EncodeToString(checksum), nil
}