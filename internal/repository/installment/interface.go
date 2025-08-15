package installment

import (
	"context"
	"finance/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, installment *entity.Installment) (*entity.Installment, error)
}
