package auth

import (
	"github.com/tapiaw38/resources-api/models"
)

// Login is function to return the user
func Login(email string, password string) (models.User, bool) {

	user, find := CheckUser(email)

	if !find {
		return user, false
	}

	if !user.PasswordMatch(password) {
		return user, false
	}

	return user, true
}
