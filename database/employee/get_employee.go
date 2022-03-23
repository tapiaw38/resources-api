package employee

import (
	"context"

	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/models"
)

// Get all employees from database
func GetEmployees(ctx context.Context) ([]models.EmployeeResponse, error) {

	q := `
	SELECT employee.id, employee.file_code, employee.agent_number, 
		employee.first_name, employee.last_name, employee.document_number, 
		employee.birth_date, employee.date_admission, employee.phone, 
		employee.address, employee.picture, employee.salary, employee.category, 
		employee.status, employee.work_number, workplace.id, workplace.name, workplace.code, 
		workplace.address, employee.created_at, employee.updated_at
		FROM employee
		LEFT JOIN workplace ON employee.workplace = workplace.id;
	`

	rows, err := database.Data().QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employees := []models.EmployeeResponse{}

	for rows.Next() {
		var e models.EmployeeResponse
		var fileCode, agentNumber, documentNumber models.NullString
		var workplaceId models.NullInt64
		var workplaceName, workplaceCode, workplaceAddress models.NullString
		var updatedAt models.NullTime
		var birthDate, dateAdmission models.NullDate

		err := rows.Scan(
			&e.ID,
			&fileCode,
			&agentNumber,
			&e.FirstName,
			&e.LastName,
			&documentNumber,
			&birthDate,
			&dateAdmission,
			&e.Phone,
			&e.Address,
			&e.Picture,
			&e.Salary,
			&e.Category,
			&e.Status,
			&e.WorkNumber,
			&workplaceId,
			&workplaceName,
			&workplaceCode,
			&workplaceAddress,
			&e.CreatedAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Set nullable fields

		if fileCode.Valid {
			e.FileCode = fileCode.String
		}

		if agentNumber.Valid {
			e.AgentNumber = agentNumber.String
		}

		if documentNumber.Valid {
			e.DocumentNumber = documentNumber.String
		}

		if birthDate.Valid {
			e.BirthDate = birthDate.String
		}

		if dateAdmission.Valid {
			e.DateAdmission = dateAdmission.String
		}

		if workplaceId.Valid {
			e.Workplace.ID = workplaceId.Int64
		}

		if workplaceName.Valid {
			e.Workplace.Name = workplaceName.String
		}

		if workplaceCode.Valid {
			e.Workplace.Code = workplaceCode.String
		}

		if workplaceAddress.Valid {
			e.Workplace.Address = workplaceAddress.String
		}

		if updatedAt.Valid {
			e.UpdatedAt = updatedAt.Time
		}

		employees = append(employees, e)
	}

	return employees, nil
}
