package models

// Employee struct
type Employee struct {
	ID             int64   `json:"id"`
	FileCode       string  `json:"file_code,omitempty"`
	AgentNumber    string  `json:"agent_number,omitempty"`
	FirstName      string  `json:"first_name,omitempty"`
	LastName       string  `json:"last_name,omitempty"`
	DocumentNumber string  `json:"document_number,omitempty"`
	BirthDate      string  `json:"birth_date,omitempty"`
	DateAdmission  string  `json:"date_admission,omitempty"`
	Phone          string  `json:"phone,omitempty"`
	Address        string  `json:"address,omitempty"`
	Picture        string  `json:"picture,omitempty"`
	Salary         float64 `json:"salary,omitempty"`
	Category       int64   `json:"category,omitempty"`
	Status         string  `json:"status,omitempty"`
	WorkNumber     string  `json:"work_number,omitempty"`
	EmployeeType   int64   `json:"employee_type,omitempty"`
	Workplace      int64   `json:"workplace,omitempty"`
	Base
}

type EmployeeResponse struct {
	ID             int64        `json:"id"`
	FileCode       string       `json:"file_code,omitempty"`
	AgentNumber    string       `json:"agent_number,omitempty"`
	FirstName      string       `json:"first_name,omitempty"`
	LastName       string       `json:"last_name,omitempty"`
	DocumentNumber string       `json:"document_number,omitempty"`
	BirthDate      string       `json:"birth_date,omitempty"`
	DateAdmission  string       `json:"date_admission,omitempty"`
	Phone          string       `json:"phone,omitempty"`
	Address        string       `json:"address,omitempty"`
	Picture        string       `json:"picture,omitempty"`
	Salary         float64      `json:"salary,omitempty"`
	Category       int64        `json:"category,omitempty"`
	Status         string       `json:"status,omitempty"`
	WorkNumber     string       `json:"work_number,omitempty"`
	EmployeeType   EmployeeType `json:"employee_type,omitempty"`
	Workplace      Workplace    `json:"workplace,omitempty"`
	Base
}
