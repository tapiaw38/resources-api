package models

// Login is the login model
type LoginResponse struct {
	User
	Token string `json:"token,omitempty"`
}
