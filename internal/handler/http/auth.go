package http

import (
	"finance/internal/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc auth.UseCase
}

func NewAuthHandler(uc auth.UseCase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) Authenticate(c *gin.Context) {
	var input auth.AuthenticateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.uc.AuthenticateWithGoogle(c.Request.Context(), input)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":         user,
		"access_token": token,
	})
}
