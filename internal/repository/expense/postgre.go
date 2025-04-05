package expense

import (
	"context"
	domain "finance/internal/domain/dto"
	"finance/internal/domain/entity"
	"fmt"
	"strconv"

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
	expenseID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Expense{}).
			Where("id = ?", id).
			Updates(expense).Error; err != nil {
			return err
		}

		expenseToUpdate := &entity.Expense{ID: uint(expenseID)}

		if err := tx.Model(expenseToUpdate).
			Association("Tags").
			Replace(expense.Tags); err != nil {
			return err
		}

		return nil
	})
}

func (r *postgresRepository) UpdateBatch(ctx context.Context, expense *entity.Expense, ids []string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		expenseIDs := make([]uint, 0, len(ids))
		for _, id := range ids {
			parsedID, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				return err
			}
			expenseIDs = append(expenseIDs, uint(parsedID))
		}

		if err := tx.Model(&entity.Expense{}).
			Where("id IN ?", ids).
			Updates(expense).Error; err != nil {
			return err
		}

		for _, expenseID := range expenseIDs {
			target := &entity.Expense{ID: expenseID}

			if err := tx.Model(target).Association("Tags").Clear(); err != nil {
				return err
			}

			if len(expense.Tags) > 0 {
				if err := tx.Model(target).Association("Tags").Append(expense.Tags); err != nil {
					return err
				}
			}
		}

		return nil
	})
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

func (r *postgresRepository) FindByFilters(ctx context.Context, filters domain.ExpenseFilters) ([]entity.Expense, int, int, error) {
	var expenses []entity.Expense
	query := r.db.WithContext(ctx).Session(&gorm.Session{})

	if filters.TimestampStart != "" {
		query = query.Where("timestamp >= ?", filters.TimestampStart)
	}

	if filters.TimestampEnd != "" {
		query = query.Where("timestamp <= ?", filters.TimestampEnd)
	}

	if filters.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filters.Name+"%")
	}

	if filters.Category != "" {
		query = query.Where("category_id = ?", filters.Category)
	}

	var sum int64
	err := query.Session(&gorm.Session{}).Model(&entity.Expense{}).Select("COALESCE(SUM(value), 0) AS sum").Scan(&sum).Error
	if err != nil {
		return nil, 0, 0, err
	}

	var total int64
	if err := query.Model(&entity.Expense{}).Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	if filters.OrderBy != "" && filters.OrderDirection != "" {
		query = query.Order(filters.OrderBy + " " + filters.OrderDirection)
	}

	offset := (filters.Page - 1) * filters.PageSize
	if err := query.Offset(offset).Limit(filters.PageSize).Preload("Tags").Preload("Category").Find(&expenses).Error; err != nil {
		return nil, 0, 0, err
	}

	return expenses, int(total), int(sum), nil
}

func (r *postgresRepository) GroupByCategory(ctx context.Context, filters domain.ExpenseFilters) ([]entity.ExpenseByGroup, error) {
	var groups []entity.ExpenseByGroup

	query := r.db.WithContext(ctx).Table("expenses AS e").
		Select("e.category_id, c.name AS category_name, c.color AS category_color, SUM(e.value) AS total_value").
		Joins("JOIN categories c ON e.category_id = c.id")

	if filters.TimestampStart != "" {
		query = query.Where("e.timestamp >= ?", filters.TimestampStart)
	}

	if filters.TimestampEnd != "" {
		query = query.Where("e.timestamp <= ?", filters.TimestampEnd)
	}

	if filters.Category != "" {
		query = query.Where("e.category_id = ?", filters.Category)
	}

	if filters.Name != "" {
		query = query.Where("e.name ILIKE ?", "%"+filters.Name+"%")
	}

	if len(filters.TagIds) > 0 {
		query = query.Where("et.tag_id IN ?", filters.TagIds)
	}

	query = query.Group("e.category_id, c.name, c.color")

	err := query.Scan(&groups).Error
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *postgresRepository) GroupByDate(
	ctx context.Context,
	filters domain.ExpenseFilters,
	unit string,
) ([]entity.ExpenseByDate, error) {
	var results []entity.ExpenseByDate

	switch unit {
	case "DAY", "MONTH", "YEAR":
	default:
		return nil, fmt.Errorf("invalid group unit: %s", unit)
	}

	groupByRule := fmt.Sprintf("EXTRACT(%s FROM e.timestamp)", unit)

	query := r.db.WithContext(ctx).
		Table("expenses AS e").
		Select(fmt.Sprintf("%s AS timestamp, c.name AS category_name, c.color AS category_color, SUM(e.value) AS total_value", groupByRule)).
		Joins("JOIN categories c ON e.category_id = c.id")

	if filters.TimestampStart != "" {
		query = query.Where("e.timestamp >= ?", filters.TimestampStart)
	}

	if filters.TimestampEnd != "" {
		query = query.Where("e.timestamp <= ?", filters.TimestampEnd)
	}

	if filters.Category != "" {
		query = query.Where("e.category_id = ?", filters.Category)
	}

	if filters.Name != "" {
		query = query.Where("e.name ILIKE ?", "%"+filters.Name+"%")
	}

	query = query.
		Group(fmt.Sprintf("c.name, c.color, %s", groupByRule)).
		Order(fmt.Sprintf("%s ASC", groupByRule))

	if err := query.Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
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
