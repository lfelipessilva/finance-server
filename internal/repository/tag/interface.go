package tag

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
)

type Repository interface {
	FindAll(ctx context.Context, filers domain.TagFilters) ([]entity.Tag, error)
}
