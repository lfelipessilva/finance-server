package http

import (
	"finance/internal/usecase/auth"
	"finance/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc     user.UseCase
	authuc auth.UseCase
}

func NewUserHandler(uc user.UseCase, authuc auth.UseCase) *UserHandler {
	return &UserHandler{uc: uc, authuc: authuc}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	var email = c.Query("email")

	users, err := h.uc.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input user.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uc.CreateUser(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.authuc.GenerateToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":         user,
		"access_token": accessToken,
	})
}
