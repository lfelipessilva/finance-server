package expense

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/repository/expense"
)

type expenseUseCase struct {
	repo expense.Repository
}

func NewExpenseUseCase(repo expense.Repository) UseCase {
	return &expenseUseCase{repo: repo}
}

func (uc *expenseUseCase) CreateExpense(ctx context.Context, input CreateExpenseInput) (*entity.Expense, error) {
	expense := &entity.Expense{
		Name:      input.Name,
		Category:  input.Category,
		Timestamp: input.Timestamp,
		Value:     input.Value,
	}

	if err := expense.Validate(); err != nil {
		return nil, err
	}

	if err := uc.repo.Create(ctx, expense); err != nil {
		return nil, err
	}

	return expense, nil
}

func (uc *expenseUseCase) GetExpenses(ctx context.Context, filters ExpenseFilters) ([]entity.Expense, error) {
	return uc.repo.FindByFilters(ctx, filters.Category, filters.MonthYear)
}
