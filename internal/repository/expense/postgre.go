package expense

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

func (r *postgresRepository) Create(ctx context.Context, expense *entity.Expense) error {
	return r.db.WithContext(ctx).Create(expense).Error
}

func (r *postgresRepository) Update(ctx context.Context, expense *entity.Expense, id string) error {
	return r.db.WithContext(ctx).
		Model(&entity.Expense{}).
		Where("id = ?", id).
		Updates(expense).Error
}

func (r *postgresRepository) CreateBatch(ctx context.Context, expenses []entity.Expense) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(expenses, 100).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *postgresRepository) FindByFilters(ctx context.Context, filters domain.ExpenseFilters) ([]entity.Expense, int, error) {
	var expenses []entity.Expense
	query := r.db.WithContext(ctx)

	if filters.TimestampStart != "" {
		query = query.Where("timestamp >= ?", filters.TimestampStart)
	}

	if filters.TimestampEnd != "" {
		query = query.Where("timestamp <= ?", filters.TimestampEnd)
	}

	var total int64
	if err := query.Model(&entity.Expense{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filters.Page - 1) * filters.PageSize
	if err := query.Offset(offset).Limit(filters.PageSize).Find(&expenses).Order("timestamp DESC").Error; err != nil {
		return nil, 0, err
	}

	return expenses, int(total), nil
}
