package secure

import "golang.org/x/crypto/bcrypt"

// CreateHash - function with create hash for password string
func CreateHash(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytePassword), err
}

// ComparePassword - function with compare hashed password and input password
func ComparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
