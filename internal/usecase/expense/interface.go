package expense

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/domain/vo"
	"time"
)

type UseCase interface {
	CreateExpense(ctx context.Context, input CreateExpenseInput) (*entity.Expense, error)
	CreateExpenses(ctx context.Context, inputs []CreateExpenseInput) ([]entity.Expense, error)
	GetExpenses(ctx context.Context, filters ExpenseFilters) ([]entity.Expense, error)
}

type CreateExpenseInput struct {
	Name      string
	Category  string
	Timestamp time.Time
	Value     float64
}

type ExpenseFilters struct {
	Category  string
	MonthYear *vo.MonthYear
}
