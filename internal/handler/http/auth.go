package http

import (
	"context"
	"finance/internal/usecase/auth"

	"google.golang.org/api/idtoken"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc auth.UseCase
}

func NewAuthHandler(uc auth.UseCase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) Authenticate(c *gin.Context) {
	// verifyGoogleToken()
	// var email = c.Query("email")

	// users, err := h.uc.Register()

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, users)
}

func verifyGoogleToken(ctx context.Context, idToken string, clientId string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(ctx, idToken, clientId)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
