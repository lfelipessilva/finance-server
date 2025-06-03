package user

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/repository/user"
)

type userUseCase struct {
	repo user.Repository
}

func NewUserUseCse(repo user.Repository) UseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return uc.repo.FindByEmail(ctx, email)
}
