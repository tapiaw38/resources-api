package user

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
)

// DeleteUser deletes a user from the database
func DeleteUser(ctx context.Context, id string) error {

	q := `
	UPDATE users
		SET is_active = false
		WHERE id = $1;
	`

	rows, err := database.Data().PrepareContext(ctx, q)

	if err != nil {
		return err
	}

	defer rows.Close()

	_, err = rows.ExecContext(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
