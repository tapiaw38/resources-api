package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim is the custom claim
type Claim struct {
	ID    uint   `json:"id,omitempty"`
	Email string `json:"email"`
	jwt.StandardClaims
}
