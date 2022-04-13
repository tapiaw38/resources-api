package employee

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// DeleteEmployee deletes an employee from the database
func DeleteEmployee(ctx context.Context, id string) (models.Employee, error) {

	var employee models.Employee

	q := `
		DELETE FROM employee
		WHERE id = $1
		RETURNING id
	`

	rows := database.Data().QueryRowContext(
		ctx, q, id,
	)

	err := rows.Scan(&employee.ID)

	if err != nil {
		return employee, err
	}

	return employee, err

}
