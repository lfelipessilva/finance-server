package user

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

func (r *postgresRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	query := r.db.WithContext(ctx).Where("email = ?", email)

	var user entity.User

	if err := query.Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
