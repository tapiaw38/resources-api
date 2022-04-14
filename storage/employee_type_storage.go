package storage

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/models"
)

// EmployeeTypeStorage is the struct that contains the database connection
type EmployeeTypeStorage struct {
	Data *Data
}

// CreateEmployeeType inserts a new employee type into the database
func (ept *EmployeeTypeStorage) CreateEmployeeType(ctx context.Context, et *models.EmployeeType) (models.EmployeeType, error) {

	var employeeType models.EmployeeType

	q := `
	INSERT INTO employee_type (name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, description, created_at, updated_at;
	`

	row := ept.Data.DB.QueryRowContext(
		ctx, q, et.Name, et.Description, time.Now(), time.Now(),
	)

	err := row.Scan(
		&employeeType.ID,
		&employeeType.Name,
		&employeeType.Description,
		&employeeType.CreatedAt,
		&employeeType.UpdatedAt,
	)

	if err != nil {
		return employeeType, err
	}

	return employeeType, nil

}

// Get all employee types from database
func (ept *EmployeeTypeStorage) GetEmployeeTypes(ctx context.Context) ([]models.EmployeeType, error) {

	q := `
	SELECT id, name, description, created_at, updated_at
		FROM employee_type;
	`

	rows, err := ept.Data.DB.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employeeTypes := []models.EmployeeType{}

	for rows.Next() {

		var et models.EmployeeType

		err := rows.Scan(
			&et.ID,
			&et.Name,
			&et.Description,
			&et.CreatedAt,
			&et.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		employeeTypes = append(employeeTypes, et)

	}

	return employeeTypes, nil

}
