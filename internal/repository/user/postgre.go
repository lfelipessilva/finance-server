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

func (r *postgresRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		First(user, user.ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}
