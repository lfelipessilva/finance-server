package http

import (
	domain "finance/internal/domain/dto"
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
	var filters domain.CategoryFilters
	filters.Name = c.Query("name")

	categories, err := h.uc.GetCategories(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
