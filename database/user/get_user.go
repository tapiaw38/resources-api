package user

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// Get all users from database
func GetUsers(ctx context.Context) ([]models.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, password, is_active, is_admin, created_at, updated_at
		FROM users;
	`

	rows, err := database.Data().QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var u models.User

		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.Email,
			&u.Picture,
			&u.PasswordHash,
			&u.IsActive,
			&u.IsAdmin,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil

}

// Get user by id from database
func GetUserById(ctx context.Context, id string) (*models.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, is_active, is_admin, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	row := database.Data().QueryRowContext(
		ctx, q, id,
	)

	u := models.User{}

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.Picture,
		&u.IsActive,
		&u.IsAdmin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil

}

// Get user by username from database
func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, is_active, is_admin, created_at, updated_at
		FROM users
		WHERE username = $1;
	`

	row := database.Data().QueryRowContext(
		ctx, q, username,
	)

	u := models.User{}

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.Picture,
		&u.IsActive,
		&u.IsAdmin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil

}
