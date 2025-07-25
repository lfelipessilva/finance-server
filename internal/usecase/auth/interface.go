package auth

import (
	"context"
	"finance/internal/domain/entity"
)

type UseCase interface {
	GenerateToken(user entity.User) (string, error)
	AuthenticateWithGoogle(ctx context.Context, input AuthenticateInput) (*entity.User, string, error)
}

type AuthenticateInput struct {
	IDToken string `json:"id_token" binding:"required"`
}
