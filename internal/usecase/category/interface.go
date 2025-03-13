package category

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
)

type UseCase interface {
	GetCategories(ctx context.Context, filters domain.CategoryFilters) ([]entity.Category, error)
}
