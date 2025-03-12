package category

import (
	"context"
	"finance/internal/domain/entity"

	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) FindAll(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category
	query := r.db.WithContext(ctx)

	if err := query.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
