package tag

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

func (r *postgresRepository) FindAll(ctx context.Context, filters domain.TagFilters) ([]entity.Tag, error) {
	var tags []entity.Tag
	query := r.db.WithContext(ctx)

	if filters.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filters.Name+"%")
	}

	if err := query.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *postgresRepository) FindById(ctx context.Context, ids []uint) ([]entity.Tag, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var tags []entity.Tag
	query := r.db.WithContext(ctx).Where("id IN (?)", ids)

	if err := query.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
