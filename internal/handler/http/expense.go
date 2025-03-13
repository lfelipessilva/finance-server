package http

import (
	domain "finance/internal/domain/dto"
	"finance/internal/usecase/expense"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	uc expense.UseCase
}

func NewExpenseHandler(uc expense.UseCase) *ExpenseHandler {
	return &ExpenseHandler{uc: uc}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var input expense.CreateExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdExpense, err := h.uc.CreateExpense(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdExpense)
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	var input expense.UpdateExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	updatedExpense, err := h.uc.UpdateExpense(c.Request.Context(), input, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, updatedExpense)
}

func (h *ExpenseHandler) CreateExpenses(c *gin.Context) {
	var inputs []expense.CreateExpenseInput

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	createdExpenses, err := h.uc.CreateExpenses(c.Request.Context(), inputs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdExpenses})
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	filters := parseFilters(c)

	expenses, total, err := h.uc.GetExpenses(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": expenses,
		"summary": gin.H{
			"total":     total,
			"page":      filters.Page,
			"page_size": filters.PageSize,
		},
	})
}

func parseFilters(c *gin.Context) domain.ExpenseFilters {
	var filters domain.ExpenseFilters

	// Parsing timestamp filters (timestamp> and timestamp<)
	if timestampStart := c.Query("timestamp_start"); timestampStart != "" {
		filters.TimestampStart = timestampStart
	}
	if timestampEnd := c.Query("timestamp_end"); timestampEnd != "" {
		filters.TimestampEnd = timestampEnd
	}

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "50")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	filters.Page = pageInt
	filters.PageSize = pageSizeInt

	return filters
}
