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

func (r *postgresRepository) Create(ctx context.Context, expense *entity.Expense) (*entity.Expense, error) {
	if err := r.db.WithContext(ctx).Create(expense).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Tags").
		First(expense, expense.ID).Error; err != nil {
		return nil, err
	}

	return expense, nil
}

func (r *postgresRepository) Update(ctx context.Context, expense *entity.Expense, id string) error {
	return r.db.WithContext(ctx).
		Model(&entity.Expense{}).
		Where("id = ?", id).
		Updates(expense).Error
}

func (r *postgresRepository) UpdateBatch(ctx context.Context, expense *entity.Expense, ids []string) error {
	return r.db.WithContext(ctx).
		Model(&entity.Expense{}).
		Where("id IN ?", ids).
		Updates(expense).Error
}

func (r *postgresRepository) CreateBatch(ctx context.Context, expenses []*entity.Expense) ([]*entity.Expense, error) {
	if err := r.db.WithContext(ctx).Create(&expenses).Error; err != nil {
		return nil, err
	}

	expenseIDs := make([]uint, len(expenses))
	for i, expense := range expenses {
		expenseIDs[i] = expense.ID
	}

	var createdExpenses []*entity.Expense
	if err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Tags").
		Find(&createdExpenses, expenseIDs).Error; err != nil {
		return nil, err
	}

	return createdExpenses, nil
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

	if filters.OrderBy != "" && filters.OrderDirection != "" {
		query = query.Order(filters.OrderBy + " " + filters.OrderDirection)
	}

	var total int64
	if err := query.Model(&entity.Expense{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filters.Page - 1) * filters.PageSize
	if err := query.Offset(offset).Limit(filters.PageSize).Preload("Tags").Preload("Category").Find(&expenses).Error; err != nil {
		return nil, 0, err
	}

	return expenses, int(total), nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).
		Model(&entity.Expense{}).
		Where("id = ?", id).
		Delete(&entity.Expense{}).Error
}

func (r *postgresRepository) DeleteBatch(ctx context.Context, ids []string) error {
	return r.db.WithContext(ctx).
		Model(&entity.Expense{}).
		Where("id IN ?", ids).
		Delete(&entity.Expense{}).Error
}
