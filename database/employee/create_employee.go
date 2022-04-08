package employee

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// CreateEmployee creates a new employee in database
func CreateEmployee(ctx context.Context, e *models.Employee) (models.Employee, error) {

	var employee models.Employee

	q := `
	INSERT INTO employee (file_code, agent_number, first_name, last_name, document_number,
		birth_date, date_admission, phone, address, picture, salary, category, work_number,
		employee_type, workplace, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		RETURNING id, file_code, agent_number, first_name, last_name, document_number,
			birth_date, date_admission, phone, address, picture, salary, category, status, work_number,
			employee_type, workplace, created_at, updated_at;
		`

	var fileCode, agentNumber, documentNumber models.NullString
	var birthDate, dateAdmission models.NullString
	var salary models.NullFloat64
	var category, employeeType, workplace models.NullInt64

	row := database.Data().QueryRowContext(
		ctx, q, database.StringToNull(e.FileCode), database.StringToNull(e.AgentNumber),
		e.FirstName, e.LastName, database.StringToNull(e.DocumentNumber),
		database.StringToNull(e.BirthDate), database.StringToNull(e.DateAdmission),
		e.Phone, e.Address, e.Picture, database.FloatToNull(e.Salary),
		database.IntToNull(e.Category), e.WorkNumber, database.IntToNull(e.EmployeeType),
		database.IntToNull(e.Workplace), e.CreatedAt, e.UpdatedAt,
	)

	err := row.Scan(
		&employee.ID,
		&fileCode,
		&agentNumber,
		&employee.FirstName,
		&employee.LastName,
		&documentNumber,
		&birthDate,
		&dateAdmission,
		&employee.Phone,
		&employee.Address,
		&employee.Picture,
		&salary,
		&category,
		&employee.Status,
		&employee.WorkNumber,
		&employeeType,
		&workplace,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	employee.FileCode = fileCode.String
	employee.AgentNumber = agentNumber.String
	employee.DocumentNumber = documentNumber.String
	employee.BirthDate = birthDate.String
	employee.DateAdmission = dateAdmission.String
	employee.Salary = salary.Float64
	employee.Category = category.Int64
	employee.EmployeeType = employeeType.Int64
	employee.Workplace = workplace.Int64

	if err != nil {
		return employee, err
	}

	return employee, nil
}
