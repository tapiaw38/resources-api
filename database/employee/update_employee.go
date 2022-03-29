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

	var fileCode, documentNumber models.NullString
	var birthDate, dateAdmission models.NullString
	var workplace models.NullInt64

	if e.FileCode == "" {
		fileCode.Valid = false
	} else {
		fileCode.Valid = true
		fileCode.String = e.FileCode
	}

	if e.DocumentNumber == "" {
		documentNumber.Valid = false
	} else {
		documentNumber.Valid = true
		documentNumber.String = e.DocumentNumber
	}

	if e.BirthDate == "" {
		birthDate.Valid = false
	} else {
		birthDate.Valid = true
		birthDate.String = e.BirthDate
	}

	if e.DateAdmission == "" {
		dateAdmission.Valid = false
	} else {
		dateAdmission.Valid = true
		dateAdmission.String = e.DateAdmission
	}

	if e.Workplace == 0 {
		workplace.Valid = false
	} else {
		workplace.Valid = true
		workplace.Int64 = e.Workplace
	}

	rows := database.Data().QueryRowContext(
		ctx, q, fileCode, e.AgentNumber, e.FirstName, e.LastName, documentNumber,
		birthDate, dateAdmission, e.Phone, e.Address, e.Picture, e.Salary,
		e.Category, e.Status, e.WorkNumber, e.EmployeeType, workplace, e.UpdatedAt, id,
	)

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

	if fileCode.Valid {
		employee.FileCode = fileCode.String
	}

	if documentNumber.Valid {
		employee.DocumentNumber = documentNumber.String
	}

	if birthDate.Valid {
		employee.BirthDate = birthDate.String
	}

	if dateAdmission.Valid {
		employee.DateAdmission = dateAdmission.String
	}

	if workplace.Valid {
		employee.Workplace = workplace.Int64
	}

	if err != nil {
		return employee, err
	}

	return employee, nil
}
