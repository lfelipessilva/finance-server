package expense

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
	"time"
)

type UseCase interface {
	CreateExpense(ctx context.Context, input CreateExpenseInput) (*entity.Expense, error)
	UpdateExpense(ctx context.Context, input UpdateExpenseInput, id string) (*entity.Expense, error)
	CreateExpenses(ctx context.Context, inputs []CreateExpenseInput) ([]*entity.Expense, error)
	GetExpenses(ctx context.Context, filters domain.ExpenseFilters) ([]entity.Expense, int, error)
}

type CreateExpenseInput struct {
	Name       string
	CategoryID uint
	TagIDs     []uint
	Bank       string
	Card       string
	Timestamp  time.Time
	Value      float64
}

type UpdateExpenseInput struct {
	Name       string
	CategoryID uint
	TagIDs     []uint
	Bank       string
	Card       string
	Timestamp  time.Time
	Value      float64
}
