package employee

import (
	"context"
)

// EmployeeStorage is the interface for the storage of employees
type Storage interface {
	CreateEmployee(ctx context.Context, employee *Employee) (Employee, error)
	GetEmployees(ctx context.Context) ([]EmployeeResponse, error)
	GetEmployeeByType(ctx context.Context, typeId string) ([]EmployeeResponse, error)
	GetEmployeeById(ctx context.Context, id string) (EmployeeResponse, error)
	UpdateEmployee(ctx context.Context, id string, employee Employee) (Employee, error)
	DeleteEmployee(ctx context.Context, id string) (Employee, error)
}
