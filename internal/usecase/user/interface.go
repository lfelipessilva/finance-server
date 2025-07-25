package user

import (
	"context"
	"finance/internal/domain/entity"
)

type UseCase interface {
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, input CreateUserInput) (*entity.User, error)
}

type CreateUserInput struct {
	Email          string
	Name           string
	Provider       string
	ProviderUserID string
	ProfilePicture string
}
