package user

import (
	"context"
	"finance/internal/domain/entity"
)

type UseCase interface {
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}
