package expense

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
	"finance/internal/repository/expense"
	"finance/internal/repository/tag"
	"fmt"
	"strconv"
	"strings"
)

type expenseUseCase struct {
	repo    expense.Repository
	tagRepo tag.Repository
}

func NewExpenseUseCase(repo expense.Repository, tagRepo tag.Repository) UseCase {
	return &expenseUseCase{repo: repo, tagRepo: tagRepo}
}

func (uc *expenseUseCase) CreateExpense(ctx context.Context, input CreateExpenseInput) (*entity.Expense, error) {
	tags, err := uc.tagRepo.FindById(ctx, input.TagIDs)
	if err != nil {
		return nil, err
	}

	expense := &entity.Expense{
		Name:         input.Name,
		Description:  input.Description,
		OriginalName: input.Name,
		Timestamp:    input.Timestamp,
		CategoryID:   input.CategoryID,
		Tags:         tags,
		Bank:         input.Bank,
		Card:         input.Card,
		Value:        input.Value,
	}

	if err := expense.Validate(); err != nil {
		return nil, err
	}

	createdExpense, err := uc.repo.Create(ctx, expense)
	if err != nil {
		return nil, err
	}

	return createdExpense, nil
}

func (uc *expenseUseCase) UpdateExpense(ctx context.Context, input UpdateExpenseInput, id string) (*entity.Expense, error) {
	tags, err := uc.tagRepo.FindById(ctx, input.TagIDs)
	if err != nil {
		return nil, err
	}
	expense := &entity.Expense{
		Name:        input.Name,
		Description: input.Description,
		Timestamp:   input.Timestamp,
		CategoryID:  input.CategoryID,
		Tags:        tags,
		Bank:        input.Bank,
		Card:        input.Card,
		Value:       input.Value,
	}

	if err := uc.repo.Update(ctx, expense, id); err != nil {
		return nil, err
	}

	return expense, nil
}

func (uc *expenseUseCase) UpdateExpenses(ctx context.Context, input UpdateExpenseInput, ids []string) ([]*entity.Expense, error) {
	tags, err := uc.tagRepo.FindById(ctx, input.TagIDs)
	if err != nil {
		return nil, err
	}

	expense := &entity.Expense{
		Name:         input.Name,
		Description:  input.Description,
		OriginalName: input.Name,
		Timestamp:    input.Timestamp,
		CategoryID:   input.CategoryID,
		Tags:         tags,
		Bank:         input.Bank,
		Card:         input.Card,
		Value:        input.Value,
	}

	var expenses []*entity.Expense
	for _, id := range ids {
		expenseID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return nil, err
		}

		expense := &entity.Expense{
			ID:           uint(expenseID),
			Name:         input.Name,
			Description:  input.Description,
			OriginalName: input.Name,
			Timestamp:    input.Timestamp,
			CategoryID:   input.CategoryID,
			Tags:         tags,
			Bank:         input.Bank,
			Card:         input.Card,
			Value:        input.Value,
		}

		expenses = append(expenses, expense)
	}

	if err := uc.repo.UpdateBatch(ctx, expense, ids); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (uc *expenseUseCase) CreateExpenses(ctx context.Context, inputs []CreateExpenseInput) ([]*entity.Expense, error) {
	uniqueTagIDs := make(map[uint]struct{})
	for _, expenseInput := range inputs {
		for _, tagID := range expenseInput.TagIDs {
			uniqueTagIDs[tagID] = struct{}{}
		}
	}

	var tagIDs []uint
	for tagID := range uniqueTagIDs {
		tagIDs = append(tagIDs, tagID)
	}

	tagIDMap := make(map[uint]entity.Tag)
	if len(tagIDs) > 0 {
		tags, err := uc.tagRepo.FindById(ctx, tagIDs)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch tags: %w", err)
		}
		for _, tag := range tags {
			tagIDMap[tag.ID] = tag
		}
	}

	var expenses []*entity.Expense
	var validationErrors []string

	for i, expenseInput := range inputs {
		var assignedTags []entity.Tag
		for _, tagID := range expenseInput.TagIDs {
			if tag, exists := tagIDMap[tagID]; exists {
				assignedTags = append(assignedTags, tag)
			}
		}

		expense := &entity.Expense{
			Name:         expenseInput.Name,
			Description:  expenseInput.Description,
			OriginalName: expenseInput.Name,
			Timestamp:    expenseInput.Timestamp,
			CategoryID:   expenseInput.CategoryID,
			Bank:         expenseInput.Bank,
			Card:         expenseInput.Card,
			Value:        expenseInput.Value,
			Tags:         assignedTags,
		}

		if err := expense.Validate(); err != nil {
			validationErrors = append(validationErrors,
				fmt.Sprintf("Invalid expense at index %d: %v", i, err))
			continue
		}

		expenses = append(expenses, expense)
	}

	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("validation errors:\n%s", strings.Join(validationErrors, "\n"))
	}

	created, err := uc.repo.CreateBatch(ctx, expenses)
	if err != nil {
		return nil, fmt.Errorf("failed to create expenses: %w", err)
	}

	return created, nil
}

func (uc *expenseUseCase) GetExpenses(ctx context.Context, filters domain.ExpenseFilters) ([]entity.Expense, int, int, error) {
	return uc.repo.FindByFilters(ctx, filters)
}

func (uc *expenseUseCase) GetExpensesByCategory(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByGroup, error) {
	return uc.repo.GroupByCategory(ctx, filters)
}

func (uc *expenseUseCase) GetExpensesByDate(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByDate, error) {
	return uc.repo.GroupByDate(ctx, filters)
}

func (uc *expenseUseCase) GetExpensesByDay(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByDate, error) {
	return uc.repo.GroupByDateUnit(ctx, filters, "DAY")
}

func (uc *expenseUseCase) GetExpensesByMonth(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByDate, error) {
	return uc.repo.GroupByDateUnit(ctx, filters, "MONTH")
}

func (uc *expenseUseCase) GetExpensesByYear(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByDate, error) {
	return uc.repo.GroupByDateUnit(ctx, filters, "YEAR")
}

func (uc *expenseUseCase) DeleteExpense(ctx context.Context, id string) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *expenseUseCase) DeleteExpenses(ctx context.Context, ids []string) error {
	err := uc.repo.DeleteBatch(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}
