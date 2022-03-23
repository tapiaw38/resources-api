package workplace

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// Get all workplaces from database
func GetWorkplaces(ctx context.Context) ([]models.Workplace, error) {

	q := `
	SELECT id, name, code, address, created_at, updated_at
		FROM workplace
		ORDER BY id ASC;
	`

	rows, err := database.Data().QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	workplaces := []models.Workplace{}

	for rows.Next() {

		var wp models.Workplace

		err := rows.Scan(
			&wp.ID,
			&wp.Name,
			&wp.Code,
			&wp.Address,
			&wp.CreatedAt,
			&wp.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		workplaces = append(workplaces, wp)

	}

	return workplaces, nil

}
