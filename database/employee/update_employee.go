package employee

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// Update employee in database by id
func UpdateEmployee(ctx context.Context, id string, e models.Employee) (models.Employee, error) {

	var employee models.Employee

	q := `
	UPDATE employee
		SET file_code = $1, agent_number = $2, first_name = $3, last_name = $4, document_number = $5,
			birth_date = $6, date_admission = $7, phone = $8, address = $9, picture = $10, salary = $11,
			category = $12, status = $13, work_number = $14, employee_type = $15, workplace = $16, updated_at = $17
		WHERE id = $18
		RETURNING id, file_code, agent_number, first_name, last_name, document_number,
			birth_date, date_admission, phone, address, picture, salary, category, status, work_number,
			employee_type, workplace, created_at, updated_at;
	`

	rows := database.Data().QueryRowContext(
		ctx, q, e.FileCode, e.AgentNumber, e.FirstName, e.LastName, e.DocumentNumber,
		e.BirthDate, e.DateAdmission, e.Phone, e.Address, e.Picture, e.Salary,
		e.Category, e.Status, e.WorkNumber, e.EmployeeType, e.Workplace, e.UpdatedAt, id,
	)

	err := rows.Scan(
		&employee.ID,
		&employee.FileCode,
		&employee.AgentNumber,
		&employee.FirstName,
		&employee.LastName,
		&employee.DocumentNumber,
		&employee.BirthDate,
		&employee.DateAdmission,
		&employee.Phone,
		&employee.Address,
		&employee.Picture,
		&employee.Salary,
		&employee.Category,
		&employee.Status,
		&employee.WorkNumber,
		&employee.EmployeeType,
		&employee.Workplace,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		return employee, err
	}

	return employee, nil
}
