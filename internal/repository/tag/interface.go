package tag

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
)

type Repository interface {
	FindAll(ctx context.Context, filters domain.TagFilters) ([]entity.Tag, error)
	FindById(ctx context.Context, ids []uint) ([]entity.Tag, error)
}
