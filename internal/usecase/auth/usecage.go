package auth

import (
	"context"
	"finance/internal/domain/entity"
	"finance/internal/usecase/user"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/api/idtoken"
)

type authUseCase struct {
	jwtSecret      string
	googleClientID string
	userUseCase    user.UseCase
}

func NewAuthUseCase(jwtSecret, googleClientID string, userUseCase user.UseCase) UseCase {
	return &authUseCase{
		jwtSecret:      jwtSecret,
		googleClientID: googleClientID,
		userUseCase:    userUseCase,
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

func (uc *authUseCase) AuthenticateWithGoogle(ctx context.Context, input AuthenticateInput) (*entity.User, string, error) {
	payload, err := idtoken.Validate(ctx, input.IDToken, uc.googleClientID)
	if err != nil {
		return nil, "", err
	}

	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)
	picture := ""
	if pic, ok := payload.Claims["picture"].(string); ok {
		picture = pic
	}
	sub := payload.Subject

	existingUser, err := uc.userUseCase.GetUserByEmail(ctx, email)

	if err == nil {
		token, err := uc.GenerateToken(existingUser)
		if err != nil {
			return nil, "", err
		}
		return &existingUser, token, nil
	}

	createInput := user.CreateUserInput{
		Email:          email,
		Name:           name,
		Provider:       "google",
		ProviderUserID: sub,
		ProfilePicture: picture,
	}

	newUser, err := uc.userUseCase.CreateUser(ctx, createInput)
	if err != nil {
		return nil, "", err
	}

	token, err := uc.GenerateToken(*newUser)
	if err != nil {
		return nil, "", err
	}

	return newUser, token, nil
}
