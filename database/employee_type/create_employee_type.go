package employeetype

import (
	"context"
	"time"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// CreateEmployeeType inserts a new employee type into the database
func CreateEmployeeType(ctx context.Context, et *models.EmployeeType) (models.EmployeeType, error) {

	var employeeType models.EmployeeType

	q := `
	INSERT INTO employee_type (name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, description, created_at, updated_at;
	`

	row := database.Data().QueryRowContext(
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
