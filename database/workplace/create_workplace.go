package workplace

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// CreateWorkplace inserts a new workplace into the database
func CreateWorkplace(ctx context.Context, w *models.Workplace) (models.Workplace, error) {

	var workplace models.Workplace

	q := `
	INSERT INTO workplace (name, code, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, code, address, created_at, updated_at;
	`

	row := database.Data().QueryRowContext(
		ctx, q, w.Name, w.Code, w.Address, time.Now(), time.Now(),
	)

	err := row.Scan(
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
