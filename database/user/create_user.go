package user

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// CreateUser inserts a new user into the database
func CreateUser(ctx context.Context, u *models.User) (models.User, error) {

	var user models.User

	q := `
    INSERT INTO users (first_name, last_name, username, email, picture, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, first_name, last_name, username, email, picture, created_at, updated_at;
    `

	if err := u.HashPassword(); err != nil {
		return user, err
	}

	row := database.Data().QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Username, u.Email,
		u.Picture, u.PasswordHash, time.Now(), time.Now(),
	)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Picture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil

}
