package expense

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, expense *entity.Expense) (*entity.Expense, error)
	Update(ctx context.Context, expense *entity.Expense, id string) error
	UpdateBatch(ctx context.Context, expense *entity.Expense, id []string) error
	CreateBatch(ctx context.Context, expense []*entity.Expense) ([]*entity.Expense, error)
	FindByFilters(ctx context.Context, filters domain.ExpenseFilters) ([]entity.Expense, int, int, error)
	GroupByCategory(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByGroup, error)
	Delete(ctx context.Context, id string) error
	DeleteBatch(ctx context.Context, ids []string) error
}
