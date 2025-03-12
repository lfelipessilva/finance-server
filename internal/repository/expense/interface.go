package expense

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/domain/vo"
)

type Repository interface {
	Create(ctx context.Context, expense *entity.Expense) error
	Update(ctx context.Context, expense *entity.Expense, id string) error
	CreateBatch(ctx context.Context, expense []entity.Expense) error
	FindByFilters(ctx context.Context, category string, monthYear *vo.MonthYear) ([]entity.Expense, error)
}
