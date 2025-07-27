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
	GetExpensesByCategory(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByGroup, error)
	GetExpensesByDate(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByDate, error)
	GetExpensesByDay(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByDate, error)
	GetExpensesByMonth(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByDate, error)
	GetExpensesByYear(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByDate, error)
	DeleteExpense(ctx context.Context, id string) error
	DeleteExpenses(ctx context.Context, ids []string) error
}

type CreateExpenseInput struct {
	UserID      uint
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
	UserID      uint
	Name        string
	CategoryID  *uint
	Description string
	TagIDs      []uint
	Bank        string
	Card        string
	Timestamp   time.Time
	Value       float64
}
