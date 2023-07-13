package utils

import "golang.org/x/crypto/bcrypt"

// This function is user to hash user password so that it can be stored in database safely
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// This function to use to compare Hash with Password to know if password is correct
func ComparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
