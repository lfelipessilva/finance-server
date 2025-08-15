package installment

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/repository/expense"
	"finance/internal/repository/installment"
	"strconv"
	"time"
)

type installmentUseCase struct {
	repo        installment.Repository
	expenseRepo expense.Repository
}

func NewInstallmentUseCase(repo installment.Repository, expenseRepo expense.Repository) UseCase {
	return &installmentUseCase{repo: repo, expenseRepo: expenseRepo}
}

// createExpenseCopy creates a copy of an expense with specific overrides
func (uc *installmentUseCase) createExpenseCopy(base entity.Expense, overrides map[string]interface{}) *entity.Expense {
	copy := &entity.Expense{
		Description:   base.Description,
		OriginalName:  base.OriginalName,
		Name:          base.Name,
		CategoryID:    base.CategoryID,
		Category:      base.Category,
		Tags:          base.Tags,
		Bank:          base.Bank,
		Card:          base.Card,
		UserID:        base.UserID,
		Timestamp:     base.Timestamp,
		Value:         base.Value,
		InstallmentID: base.InstallmentID,
	}

	// Apply overrides
	if value, ok := overrides["value"]; ok {
		copy.Value = value.(float64)
	}
	if timestamp, ok := overrides["timestamp"]; ok {
		copy.Timestamp = timestamp.(time.Time)
	}
	if installmentID, ok := overrides["installmentID"]; ok {
		copy.InstallmentID = installmentID.(*uint)
	}

	return copy
}

func (uc *installmentUseCase) CreateInstallment(ctx context.Context, input CreateInstallmentInput) (*entity.Installment, error) {
	expense, err := uc.expenseRepo.FindById(ctx, input.ExpenseId)
	if err != nil {
		return nil, err
	}

	expenses := make([]*entity.Expense, int(input.Quantity)-1)

	installmentValue := expense.Value / input.Quantity

	installment := &entity.Installment{
		Value:          expense.Value,
		Quantity:       int(input.Quantity),
		TimestampStart: expense.Timestamp,
		TimestampEnd:   expense.Timestamp.AddDate(0, int(input.Quantity), 0),
	}

	if err := installment.Validate(); err != nil {
		return nil, err
	}

	createdInstallment, err := uc.repo.Create(ctx, installment)
	if err != nil {
		return nil, err
	}

	for i := range expenses {
		expenses[i] = uc.createExpenseCopy(expense, map[string]interface{}{
			"value":         installmentValue,
			"timestamp":     expense.Timestamp.AddDate(0, i+1, 0),
			"installmentID": &createdInstallment.ID,
		})
	}

	_, err = uc.expenseRepo.CreateBatch(ctx, expenses)
	if err != nil {
		return nil, err
	}

	updatedExpense := uc.createExpenseCopy(expense, map[string]interface{}{
		"value":         installmentValue,
		"installmentID": &createdInstallment.ID,
	})
	updatedExpense.ID = input.ExpenseId

	err = uc.expenseRepo.Update(ctx, updatedExpense, strconv.Itoa(int(input.ExpenseId)))
	if err != nil {
		return nil, err
	}

	return createdInstallment, nil
}
