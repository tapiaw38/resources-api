package auth

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tapiaw38/resources-api/database/auth"
	"github.com/tapiaw38/resources-api/models"
)

var userFind models.User

// ValidateToken validates the token
func ValidateToken(token string) (*models.Claim, bool, models.User, error) {

	key := []byte(os.Getenv("JWT_SECRET"))

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	u := models.User{}

	if len(splitToken) != 2 {
		return claims, false, u, errors.New("The token format is invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		user, find := auth.CheckUser(claims.Email)

		if find {
			userFind = user
		}

		return claims, find, user, nil

	}

	if !tkn.Valid {
		return claims, false, u, errors.New("Invalid token")
	}

	return claims, false, u, err

}

// Get the user from the token
func GetUser() models.User {
	return userFind
}
