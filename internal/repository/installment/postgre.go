package installment

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

func (r *postgresRepository) Create(ctx context.Context, installment *entity.Installment) (*entity.Installment, error) {
	if err := r.db.WithContext(ctx).Create(installment).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		First(installment, installment.ID).Error; err != nil {
		return nil, err
	}

	return installment, nil
}
