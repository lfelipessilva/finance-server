package user

import (
	"context"
	"finance/internal/domain/entity"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
