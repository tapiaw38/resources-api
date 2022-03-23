package models

// User is the user model
type User struct {
	ID           uint   `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	Picture      string `json:"picture,omitempty"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"-"`
	IsActive     bool   `json:"is_active,omitempty"`
	IsAdmin      bool   `json:"is_admin,omitempty"`
	Base
}
