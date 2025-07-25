package auth

import (
	"finance/internal/domain/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

type authUseCase struct {
	jwtSecret string
}

func NewAuthUseCase(jwtSecret string) UseCase {
	return &authUseCase{
		jwtSecret: jwtSecret,
	}
}

func (uc *authUseCase) GenerateToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp":     time.Now().Add(time.Hour * 24 * 365).Unix(),
		})

	tokenString, err := token.SignedString([]byte(uc.jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
