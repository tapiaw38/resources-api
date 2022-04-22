package employee_type

import "context"

// EmployeeTypeStorage is the interface for the storage of employees types
type Storage interface {
	CreateEmployeeType(ctx context.Context, employeeType *EmployeeType) (EmployeeType, error)
	GetEmployeeTypes(ctx context.Context) ([]EmployeeType, error)
}
