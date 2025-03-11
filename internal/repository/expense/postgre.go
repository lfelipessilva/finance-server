package expense

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/domain/vo"
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

func (r *postgresRepository) FindByFilters(ctx context.Context, category string, my *vo.MonthYear) ([]entity.Expense, error) {
	var expenses []entity.Expense
	query := r.db.WithContext(ctx)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if my != nil {
		start, end := my.TimeRange()
		query = query.Where("timestamp >= ? AND timestamp < ?", start, end)
	}

	if err := query.Find(&expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}