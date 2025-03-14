package http

import (
	domain "finance/internal/domain/dto"
	"finance/internal/usecase/tag"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	uc tag.UseCase
}

func NewTagHandler(uc tag.UseCase) *TagHandler {
	return &TagHandler{uc: uc}
}

func (h *TagHandler) GetCategories(c *gin.Context) {
	var filters domain.TagFilters
	filters.Name = c.Query("name")

	tags, err := h.uc.GetTags(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}
