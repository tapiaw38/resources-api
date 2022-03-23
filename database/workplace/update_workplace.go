package workplace

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// UpdateWorkplace updates a workplace in the database
func UpdateWorkplace(ctx context.Context, id string, w models.Workplace) (models.Workplace, error) {

	var workplace models.Workplace

	q := `
	UPDATE workplace
		SET name = $1, code = $2, address = $3, updated_at = $4
		WHERE id = $5
		RETURNING id, name, code, address, created_at, updated_at;
	`

	rows := database.Data().QueryRowContext(
		ctx, q, w.Name, w.Code, w.Address, time.Now(), id,
	)

	err := rows.Scan(
		&workplace.ID,
		&workplace.Name,
		&workplace.Code,
		&workplace.Address,
		&workplace.CreatedAt,
		&workplace.UpdatedAt,
	)

	if err != nil {
		return workplace, err
	}

	return workplace, nil

}
