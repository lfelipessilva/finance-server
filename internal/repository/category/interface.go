package category

import (
	"context"
	"finance/internal/domain/entity"
)

type Repository interface {
	FindAll(ctx context.Context) ([]entity.Category, error)
}
