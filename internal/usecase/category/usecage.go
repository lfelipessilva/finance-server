package category

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/repository/category"
)

type categoryUseCase struct {
	repo category.Repository
}

func NewCategoryUseCse(repo category.Repository) UseCase {
	return &categoryUseCase{repo: repo}
}

func (uc *categoryUseCase) GetCategories(ctx context.Context) ([]entity.Category, error) {
	return uc.repo.FindAll(ctx)
}
