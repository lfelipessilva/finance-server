package tag

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
)

type UseCase interface {
	GetTags(ctx context.Context, filters domain.TagFilters) ([]entity.Tag, error)
}
