package expense

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/repository/expense"
	"fmt"
	"strings"
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

func (uc *expenseUseCase) CreateExpenses(ctx context.Context, inputs []CreateExpenseInput) ([]entity.Expense, error) {
	var expenses []entity.Expense
	var validationErrors []string

	for i, input := range inputs {
		expense := entity.Expense{
			Name:      input.Name,
			Category:  input.Category,
			Timestamp: input.Timestamp,
			Value:     input.Value,
		}

		if err := expense.Validate(); err != nil {
			validationErrors = append(validationErrors,
				fmt.Sprintf("Invalid expense at index %d: %v", i, err))
			continue
		}

		expenses = append(expenses, expense)
	}

	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("validation errors:\n%s",
			strings.Join(validationErrors, "\n"))
	}

	if err := uc.repo.CreateBatch(ctx, expenses); err != nil {
		return nil, fmt.Errorf("failed to create expenses: %w", err)
	}

	return expenses, nil
}

func (uc *expenseUseCase) GetExpenses(ctx context.Context, filters ExpenseFilters) ([]entity.Expense, error) {
	return uc.repo.FindByFilters(ctx, filters.Category, filters.MonthYear)
}
