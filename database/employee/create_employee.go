package employee

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

func CreateEmployee(ctx context.Context, e *models.Employee) (models.Employee, error) {

	var employee models.Employee

	q := `
	INSERT INTO employees (file_code, agent_number, first_name, last_name, document_number,
		birth_date, date_admission, phone, address, picture, salary, category, status, work_number,
		employee_type, workplace, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
		RETURNING id, file_code, agent_number, first_name, last_name, document_number,
			birth_date, date_admission, phone, address, picture, salary, category, status, work_number,
			employee_type, workplace, created_at, updated_at;
		`

	row := database.Data().QueryRowContext(
		ctx, q, e.FileCode, e.AgentNumber, e.FirstName, e.LastName, e.DocumentNumber,
		e.BirthDate, e.DateAdmission, e.Phone, e.Address, e.Picture, e.Salary,
		e.Category, e.Status, e.WorkNumber, e.EmployeeType, e.Workplace, e.CreatedAt, e.UpdatedAt,
	)

	err := row.Scan(
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
