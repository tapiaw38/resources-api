package workplace

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
)

func DeleteWorkplace(ctx context.Context, id string) error {

	q := `
	DELETE FROM workplace
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
