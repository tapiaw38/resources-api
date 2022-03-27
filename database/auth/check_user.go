package auth

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// CheckUser checks if a user and email exists in the database
func CheckUser(email string) (models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	q := `
	SELECT id, first_name, last_name, username, email, password, picture, is_active, is_admin, created_at, updated_at
		FROM users
		WHERE email = $1;
	`

	rows := database.Data().QueryRowContext(ctx, q, email)

	var user models.User

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Picture,
		&user.IsActive,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, false
	}

	return user, true

}
