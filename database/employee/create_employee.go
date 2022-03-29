package employee

import (
	"context"
	"log"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

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

	if e.FileCode == "" {
		fileCode.Valid = false
	} else {
		fileCode.Valid = true
		fileCode.String = e.FileCode
	}

	if e.AgentNumber == "" {
		agentNumber.Valid = false
	} else {
		agentNumber.Valid = true
		agentNumber.String = e.AgentNumber
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

	if e.Salary == 0 {
		salary.Valid = false
	} else {
		log.Println("Salary:", e.Salary, salary.Float64)
		salary.Valid = true
		salary.Float64 = float64(e.Salary)
	}

	if e.Category == 0 {
		category.Valid = false
	} else {
		category.Valid = true
		category.Int64 = e.Category
	}

	if e.EmployeeType == 0 {
		employeeType.Valid = false
	} else {
		employeeType.Valid = true
		employeeType.Int64 = int64(e.EmployeeType)
	}

	if e.Workplace == 0 {
		workplace.Valid = false
	} else {
		workplace.Valid = true
		workplace.Int64 = int64(e.Workplace)
	}

	row := database.Data().QueryRowContext(
		ctx, q, fileCode, agentNumber, e.FirstName, e.LastName, documentNumber,
		birthDate, dateAdmission, e.Phone, e.Address, e.Picture, salary,
		category, e.WorkNumber, employeeType, workplace, e.CreatedAt, e.UpdatedAt,
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

	if fileCode.Valid {
		employee.FileCode = fileCode.String
	}

	if agentNumber.Valid {
		employee.AgentNumber = agentNumber.String
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

	if salary.Valid {
		employee.Salary = salary.Float64
	}

	if category.Valid {
		employee.Category = category.Int64
	}

	if employeeType.Valid {
		employee.EmployeeType = employeeType.Int64
	}

	if workplace.Valid {
		employee.Workplace = workplace.Int64
	}

	if err != nil {
		return employee, err
	}

	return employee, nil
}
