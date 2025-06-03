package http

import (
	"finance/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc user.UseCase
}

func NewUserHandler(uc user.UseCase) *UserHandler {
	return &UserHandler{uc: uc}
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
	var email = c.Query("email")

	users, err := h.uc.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
