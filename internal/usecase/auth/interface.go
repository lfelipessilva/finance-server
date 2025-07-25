package auth

import (
	"finance/internal/domain/entity"
)

type UseCase interface {
	GenerateToken(user entity.User) (string, error)
}
