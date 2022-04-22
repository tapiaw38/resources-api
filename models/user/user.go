package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// User is the user model
type User struct {
	ID           uint      `json:"id,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Picture      string    `json:"picture,omitempty"`
	Password     string    `json:"password,omitempty"`
	PasswordHash string    `json:"-"`
	IsActive     bool      `json:"is_active,omitempty"`
	IsAdmin      bool      `json:"is_admin,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

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

// LoginResponse is the login response
type LoginResponse struct {
	User
	Token string `json:"token,omitempty"`
}

// Claim is the custom claim
type Claim struct {
	ID    uint   `json:"id,omitempty"`
	Email string `json:"email"`
	jwt.StandardClaims
}
