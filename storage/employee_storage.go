package storage

import (
	"context"

	"github.com/tapiaw38/resources-api/models"
)

type EmployeeStorage struct {
	Data *Data
}

// CreateEmployee creates a new employee in database
func (ep *EmployeeStorage) CreateEmployee(ctx context.Context, e *models.Employee) (models.Employee, error) {

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

	row := ep.Data.DB.QueryRowContext(
		ctx, q, StringToNull(e.FileCode), StringToNull(e.AgentNumber),
		e.FirstName, e.LastName, StringToNull(e.DocumentNumber),
		StringToNull(e.BirthDate), StringToNull(e.DateAdmission),
		e.Phone, e.Address, e.Picture, FloatToNull(e.Salary),
		IntToNull(e.Category), e.WorkNumber, IntToNull(e.EmployeeType),
		IntToNull(e.Workplace), e.CreatedAt, e.UpdatedAt,
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

// Get all employees from database
func (ep *EmployeeStorage) GetEmployees(ctx context.Context) ([]models.EmployeeResponse, error) {

	q := `
	SELECT employee.id, employee.file_code, employee.agent_number, 
		employee.first_name, employee.last_name, employee.document_number, 
		employee.birth_date, employee.date_admission, employee.phone, 
		employee.address, employee.picture, employee.salary, employee.category, 
		employee.status, employee.work_number, workplace.id, workplace.name, workplace.code, 
		workplace.address, employee_type.id, employee_type.name, employee_type.description, 
		employee.created_at, employee.updated_at	
		FROM employee
		LEFT JOIN workplace ON employee.workplace = workplace.id
		LEFT JOIN employee_type ON employee.employee_type = employee_type.id;
	`

	rows, err := ep.Data.DB.QueryContext(ctx, q)
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
		var typeId models.NullInt64
		var typeName, typeDescription models.NullString
		var birthDate, dateAdmission models.NullString

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
			&typeId,
			&typeName,
			&typeDescription,
			&e.CreatedAt,
			&e.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Set nullable fields

		e.FileCode = fileCode.String
		e.AgentNumber = agentNumber.String
		e.DocumentNumber = documentNumber.String
		e.BirthDate = birthDate.String
		e.DateAdmission = dateAdmission.String
		e.Workplace.ID = workplaceId.Int64
		e.Workplace.Name = workplaceName.String
		e.Workplace.Code = workplaceCode.String
		e.Workplace.Address = workplaceAddress.String
		e.EmployeeType.ID = typeId.Int64
		e.EmployeeType.Name = typeName.String
		e.EmployeeType.Description = typeDescription.String

		employees = append(employees, e)
	}

	return employees, nil
}

// Get employees by type from database
func (ep *EmployeeStorage) GetEmployeeByType(ctx context.Context, typeId string) ([]models.EmployeeResponse, error) {

	q := `
	SELECT employee.id, employee.file_code, employee.agent_number, 
		employee.first_name, employee.last_name, employee.document_number, 
		employee.birth_date, employee.date_admission, employee.phone, 
		employee.address, employee.picture, employee.salary, employee.category, 
		employee.status, employee.work_number, workplace.id, workplace.name, workplace.code, 
		workplace.address, employee_type.id, employee_type.name, employee_type.description, 
		employee.created_at, employee.updated_at	
		FROM employee
		LEFT JOIN workplace ON employee.workplace = workplace.id
		LEFT JOIN employee_type ON employee.employee_type = employee_type.id
		WHERE employee.employee_type = $1;
	`

	rows, err := ep.Data.DB.QueryContext(ctx, q, typeId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employees := []models.EmployeeResponse{}

	for rows.Next() {
		var e models.EmployeeResponse
		var fileCode, agentNumber, documentNumber models.NullString
		var birthDate, dateAdmission models.NullString
		var salary models.NullFloat64
		var category models.NullInt64
		// var workplace nulls
		var workplaceId models.NullInt64
		var workplaceName, workplaceCode, workplaceAddress models.NullString
		// var type nulls
		var typeId models.NullInt64
		var typeName, typeDescription models.NullString

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
			&salary,
			&category,
			&e.Status,
			&e.WorkNumber,
			&workplaceId,
			&workplaceName,
			&workplaceCode,
			&workplaceAddress,
			&typeId,
			&typeName,
			&typeDescription,
			&e.CreatedAt,
			&e.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Set nullable fields
		e.FileCode = fileCode.String
		e.AgentNumber = agentNumber.String
		e.DocumentNumber = documentNumber.String
		e.BirthDate = birthDate.String
		e.DateAdmission = dateAdmission.String
		e.Salary = salary.Float64
		e.Category = category.Int64
		// Set workplace
		e.Workplace.ID = workplaceId.Int64
		e.Workplace.Name = workplaceName.String
		e.Workplace.Code = workplaceCode.String
		e.Workplace.Address = workplaceAddress.String
		// Set employee type
		e.EmployeeType.ID = typeId.Int64
		e.EmployeeType.Name = typeName.String
		e.EmployeeType.Description = typeDescription.String

		employees = append(employees, e)
	}

	return employees, nil
}

// Get employees by id from database
func (ep *EmployeeStorage) GetEmployeeById(ctx context.Context, id string) (models.EmployeeResponse, error) {

	q := `
	SELECT employee.id, employee.file_code, employee.agent_number, 
		employee.first_name, employee.last_name, employee.document_number, 
		employee.birth_date, employee.date_admission, employee.phone, 
		employee.address, employee.picture, employee.salary, employee.category, 
		employee.status, employee.work_number, workplace.id, workplace.name, workplace.code, 
		workplace.address, employee_type.id, employee_type.name, employee_type.description, 
		employee.created_at, employee.updated_at	
		FROM employee
		LEFT JOIN workplace ON employee.workplace = workplace.id
		LEFT JOIN employee_type ON employee.employee_type = employee_type.id
		WHERE employee.id = $1;
	`

	row := ep.Data.DB.QueryRowContext(ctx, q, id)

	var e models.EmployeeResponse
	var fileCode, agentNumber, documentNumber models.NullString
	var birthDate, dateAdmission models.NullString
	var salary models.NullFloat64
	var category models.NullInt64
	// var workplace nulls
	var workplaceId models.NullInt64
	var workplaceName, workplaceCode, workplaceAddress models.NullString
	// var type nulls
	var typeId models.NullInt64
	var typeName, typeDescription models.NullString

	err := row.Scan(
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
		&salary,
		&category,
		&e.Status,
		&e.WorkNumber,
		&workplaceId,
		&workplaceName,
		&workplaceCode,
		&workplaceAddress,
		&typeId,
		&typeName,
		&typeDescription,
		&e.CreatedAt,
		&e.UpdatedAt,
	)

	if err != nil {
		return e, err
	}

	// Set nullable fields
	e.FileCode = fileCode.String
	e.AgentNumber = agentNumber.String
	e.DocumentNumber = documentNumber.String
	e.BirthDate = birthDate.String
	e.DateAdmission = dateAdmission.String
	e.Salary = salary.Float64
	e.Category = category.Int64
	// Set workplace
	e.Workplace.ID = workplaceId.Int64
	e.Workplace.Name = workplaceName.String
	e.Workplace.Code = workplaceCode.String
	e.Workplace.Address = workplaceAddress.String
	// Set employee type
	e.EmployeeType.ID = typeId.Int64
	e.EmployeeType.Name = typeName.String
	e.EmployeeType.Description = typeDescription.String

	return e, nil
}

// Update employee in database by id
func (ep *EmployeeStorage) UpdateEmployee(ctx context.Context, id string, e models.Employee) (models.Employee, error) {

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

	rows := ep.Data.DB.QueryRowContext(
		ctx, q, StringToNull(e.FileCode), e.AgentNumber, e.FirstName,
		e.LastName, StringToNull(e.DocumentNumber),
		StringToNull(e.BirthDate), StringToNull(e.DateAdmission),
		e.Phone, e.Address, e.Picture, e.Salary, e.Category, e.Status, e.WorkNumber,
		e.EmployeeType, IntToNull(e.Workplace), e.UpdatedAt, id,
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

	employee.FileCode = fileCode.String
	employee.DocumentNumber = documentNumber.String
	employee.BirthDate = birthDate.String
	employee.DateAdmission = dateAdmission.String
	employee.Workplace = workplace.Int64

	if err != nil {
		return employee, err
	}

	return employee, nil
}

// DeleteEmployee deletes an employee from the database
func (ep *UserStorage) DeleteEmployee(ctx context.Context, id string) (models.Employee, error) {

	var employee models.Employee

	q := `
		DELETE FROM employee
		WHERE id = $1
		RETURNING id
	`

	rows := ep.Data.DB.QueryRowContext(
		ctx, q, id,
	)

	err := rows.Scan(&employee.ID)

	if err != nil {
		return employee, err
	}

	return employee, err

}
