package util

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	lowerCharSet   = "abcdedfghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = upperCharSet + lowerCharSet + specialCharSet + numberSet
)

// GeneratePassword generates a random password
func GeneratePassword(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	password := make([]byte, length)

	// Aseguramos que la contraseña tenga al menos un caracter en mayúscula, uno en minúscula, uno numérico y uno especial
	password[0] = allCharSet[rand.Intn(26)]    // al menos una letra mayúscula
	password[1] = allCharSet[rand.Intn(26)+26] // al menos una letra minúscula
	password[2] = allCharSet[rand.Intn(10)+52] // al menos un número
	password[3] = allCharSet[rand.Intn(7)+62]  // al menos un caracter especial

	for i := 4; i < length; i++ {
		password[i] = allCharSet[rand.Intn(len(allCharSet))]
	}

	for i := range password {
		j := rand.Intn(i + 1)
		password[i], password[j] = password[j], password[i]
	}

	return string(password)
}

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	return string(bytes), err
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return (err == nil)
}
