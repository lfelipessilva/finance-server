package expense

import (
	"context"
	domain "finance/internal/domain/dto"
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
	fmt.Println(input)
	expense := &entity.Expense{
		Name:       input.Name,
		Timestamp:  input.Timestamp,
		CategoryID: input.CategoryID,
		Bank:       input.Bank,
		Card:       input.Card,
		Value:      input.Value,
	}

	if err := expense.Validate(); err != nil {
		return nil, err
	}

	expense, err := uc.repo.Create(ctx, expense)

	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (uc *expenseUseCase) UpdateExpense(ctx context.Context, input UpdateExpenseInput, id string) (*entity.Expense, error) {
	expense := &entity.Expense{
		Name:       input.Name,
		Timestamp:  input.Timestamp,
		CategoryID: input.CategoryID,
		Bank:       input.Bank,
		Card:       input.Card,
		Value:      input.Value,
	}

	if err := uc.repo.Update(ctx, expense, id); err != nil {
		return nil, err
	}

	return expense, nil
}

func (uc *expenseUseCase) CreateExpenses(ctx context.Context, inputs []CreateExpenseInput) ([]*entity.Expense, error) {
	var expenses []entity.Expense
	var validationErrors []string

	for i, input := range inputs {
		expense := entity.Expense{
			Name:       input.Name,
			Timestamp:  input.Timestamp,
			CategoryID: input.CategoryID,
			Bank:       input.Bank,
			Card:       input.Card,
			Value:      input.Value,
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

	created, err := uc.repo.CreateBatch(ctx, expenses)

	if err != nil {
		return nil, fmt.Errorf("failed to create expenses: %w", err)
	}

	return created, nil
}

func (uc *expenseUseCase) GetExpenses(ctx context.Context, filters domain.ExpenseFilters) ([]entity.Expense, int, error) {
	return uc.repo.FindByFilters(ctx, filters)
}
