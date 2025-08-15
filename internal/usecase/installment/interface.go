package installment

import (
	"context"
	"finance/internal/domain/entity"
)

type UseCase interface {
	CreateInstallment(ctx context.Context, input CreateInstallmentInput) (*entity.Installment, error)
}

type CreateInstallmentInput struct {
	ExpenseId uint    `json:"expense_id"`
	Quantity  float64 `json:"quantity"`
}
