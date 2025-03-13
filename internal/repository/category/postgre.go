package category

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"

	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) FindAll(ctx context.Context, filters domain.CategoryFilters) ([]entity.Category, error) {
	var categories []entity.Category
	query := r.db.WithContext(ctx)

	if filters.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filters.Name+"%")
	}

	if err := query.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
