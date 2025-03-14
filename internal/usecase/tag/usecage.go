package tag

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
	"finance/internal/repository/tag"
)

type tagUseCase struct {
	repo tag.Repository
}

func NewTagUseCse(repo tag.Repository) UseCase {
	return &tagUseCase{repo: repo}
}

func (uc *tagUseCase) GetTags(ctx context.Context, filters domain.TagFilters) ([]entity.Tag, error) {
	return uc.repo.FindAll(ctx, filters)
}
