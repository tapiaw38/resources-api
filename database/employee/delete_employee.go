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
		RETURNING id, file_code, agent_number, first_name, last_name, document_number,
			birth_date, date_admission, phone, address, picture, salary, category, status, work_number,
			employee_type, workplace, created_at, updated_at;
	`

	rows := database.Data().QueryRowContext(
		ctx, q, id,
	)

	var fileCode, documentNumber models.NullString
	var birthDate, dateAdmission models.NullString
	var workplace models.NullInt64

	err := rows.Scan(
		&employee.ID,
		&fileCode,
		&employee.AgentNumber,
		&employee.FirstName,
		&employee.LastName,
		&documentNumber,
		&birthDate,
		&dateAdmission,
		&employee.Phone,
		&employee.Address,
		&employee.Picture,
		&employee.Salary,
		&employee.Category,
		&employee.Status,
		&employee.WorkNumber,
		&employee.EmployeeType,
		&workplace,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	employee.FileCode = fileCode.String
	employee.DocumentNumber = documentNumber.String
	employee.BirthDate = birthDate.String
	employee.DateAdmission = dateAdmission.String
	employee.Workplace = workplace.Int64

	return employee, err
}
