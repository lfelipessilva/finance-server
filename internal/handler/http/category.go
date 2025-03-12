package http

import (
	"finance/internal/usecase/category"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	uc category.UseCase
}

func NewCategoryHandler(uc category.UseCase) *CategoryHandler {
	return &CategoryHandler{uc: uc}
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	expenses, err := h.uc.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}
