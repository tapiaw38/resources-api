package storage

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/models"
)

type WorkplaceStorage struct {
	Data *Data
}

// CreateWorkplace inserts a new workplace into the database
func (wp *WorkplaceStorage) CreateWorkplace(ctx context.Context, w *models.Workplace) (models.Workplace, error) {

	var workplace models.Workplace

	q := `
	INSERT INTO workplace (name, code, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, code, address, created_at, updated_at;
	`

	row := wp.Data.DB.QueryRowContext(
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

// Get all workplaces from database
func (wp *WorkplaceStorage) GetWorkplaces(ctx context.Context) ([]models.Workplace, error) {

	q := `
	SELECT id, name, code, address, created_at, updated_at
		FROM workplace
		ORDER BY id ASC;
	`

	rows, err := wp.Data.DB.QueryContext(ctx, q)

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

// UpdateWorkplace updates a workplace in the database
func (wp *WorkplaceStorage) UpdateWorkplace(ctx context.Context, id string, w models.Workplace) (models.Workplace, error) {

	var workplace models.Workplace

	q := `
	UPDATE workplace
		SET name = $1, code = $2, address = $3, updated_at = $4
		WHERE id = $5
		RETURNING id, name, code, address, created_at, updated_at;
	`

	rows := wp.Data.DB.QueryRowContext(
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

func (wp *WorkplaceStorage) DeleteWorkplace(ctx context.Context, id string) error {

	q := `
	DELETE FROM workplace
		WHERE id = $1;
	`

	rows, err := wp.Data.DB.PrepareContext(ctx, q)

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
