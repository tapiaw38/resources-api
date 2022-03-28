package auth

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	auth "github.com/tapiaw38/resources-api/database/auth"
	"github.com/tapiaw38/resources-api/models"
)

var userFind models.User

// ValidateToken validates the token
func ValidateToken(tk string) (*models.Claim, bool, models.User, error) {

	u := models.User{}

	myKey := []byte(os.Getenv("JWT_SECRET"))

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, u, errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		user, find := auth.CheckUser(claims.Email)

		if find {
			userFind = user
		}

		return claims, find, user, nil

	}

	if !tkn.Valid {
		return claims, false, u, errors.New("invalid token")
	}

	return claims, false, u, err

}

// Get the user from the token
func GetUser() models.User {
	return userFind
}
