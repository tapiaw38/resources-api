package user

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// UpdateUser updates a user in the database
func UpdateUser(ctx context.Context, id string, u models.User) (models.User, error) {

	var user models.User

	q := `
	UPDATE users
		SET first_name = $1, last_name = $2, username = $3, email = $4, picture = $5, is_active = $6, is_admin = $7, updated_at = $8
		WHERE id = $9
		RETURNING id, first_name, last_name, username, email, picture, is_active, is_admin, created_at, updated_at;
	`

	rows := database.Data().QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Username, u.Email,
		u.Picture, u.IsActive, u.IsAdmin, time.Now(), id,
	)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Picture,
		&user.IsActive,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil

}
