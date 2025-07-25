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

func (uc *userUseCase) CreateUser(ctx context.Context, input CreateUserInput) (*entity.User, error) {
	user := &entity.User{
		Name:           input.Name,
		Email:          input.Email,
		Provider:       input.Provider,
		ProviderUserID: input.ProviderUserID,
		ProfilePicture: input.ProfilePicture,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	createdUser, err := uc.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
