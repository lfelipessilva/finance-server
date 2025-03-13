package category

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
	"finance/internal/repository/category"
)

type categoryUseCase struct {
	repo category.Repository
}

func NewCategoryUseCse(repo category.Repository) UseCase {
	return &categoryUseCase{repo: repo}
}

func (uc *categoryUseCase) GetCategories(ctx context.Context, filters domain.CategoryFilters) ([]entity.Category, error) {
	return uc.repo.FindAll(ctx, filters)
}
