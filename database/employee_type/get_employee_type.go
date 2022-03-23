package employeetype

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// Get all employee types from database
func GetEmployeeTypes(ctx context.Context) ([]models.EmployeeType, error) {

	q := `
	SELECT id, name, description, created_at, updated_at
		FROM employee_type;
	`

	rows, err := database.Data().QueryContext(ctx, q)

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
