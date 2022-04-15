package workplace

import "context"

// WorkplaceStorage is the interface for the storage of workplaces
type Storage interface {
	CreateWorkplace(ctx context.Context, workplace *Workplace) (Workplace, error)
	GetWorkplaces(ctx context.Context) ([]Workplace, error)
	UpdateWorkplace(ctx context.Context, id string, workplace Workplace) (Workplace, error)
	DeleteWorkplace(ctx context.Context, id string) error
}
