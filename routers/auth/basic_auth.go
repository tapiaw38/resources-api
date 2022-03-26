package auth

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	login "github.com/tapiaw38/resources-api/database/login"
	"github.com/tapiaw38/resources-api/models"
)

var UserEmail string
var UserID uint

// ParseClaims parses the token and returns the claims
func BasicAuth(token string) (*models.Claim, bool, uint, error) {

	key := []byte(os.Getenv("JWT_SECRET"))

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, 0, errors.New("Invalid token")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		_, find := login.CheckUser(claims.Email)

		if find {
			UserEmail = claims.Email
			UserID = claims.ID
		}

		return claims, find, UserID, nil

	}

	if !tkn.Valid {
		return claims, false, 0, errors.New("Invalid token")
	}

	return claims, false, 0, err

}
