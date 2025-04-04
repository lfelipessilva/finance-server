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
	UpdateExpenses(ctx context.Context, input UpdateExpenseInput, ids []string) ([]*entity.Expense, error)
	CreateExpenses(ctx context.Context, inputs []CreateExpenseInput) ([]*entity.Expense, error)
	GetExpenses(ctx context.Context, filters domain.ExpenseFilters) ([]entity.Expense, int, int, error)
	GetExpensesByGroup(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByGroup, error)
	DeleteExpense(ctx context.Context, id string) error
	DeleteExpenses(ctx context.Context, ids []string) error
}

type CreateExpenseInput struct {
	Name        string
	CategoryID  *uint
	Description string
	TagIDs      []uint
	Bank        string
	Card        string
	Timestamp   time.Time
	Value       float64
}

type UpdateExpenseInput struct {
	Name        string
	CategoryID  *uint
	Description string
	TagIDs      []uint
	Bank        string
	Card        string
	Timestamp   time.Time
	Value       float64
}
