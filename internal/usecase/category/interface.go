package category

import (
	"context"
	"finance/internal/domain/entity"
)

type UseCase interface {
	GetCategories(ctx context.Context) ([]entity.Category, error)
}
