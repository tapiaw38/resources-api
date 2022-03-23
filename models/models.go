package models

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the password
func (u *User) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(passwordHash)

	return nil
}

// CheckPassword checks the password
func (u *User) PasswordMatch(password string) bool {
	passwordBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), passwordBytes)

	return err == nil
}
