package http

import (
	domain "finance/internal/domain/dto"
	"finance/internal/usecase/expense"
	"fmt"
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
	userID, _ := c.Get("user_id")
	var input expense.CreateExpenseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = userID.(uint)

	createdExpense, err := h.uc.CreateExpense(c.Request.Context(), input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdExpense)
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var input expense.UpdateExpenseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = userID.(uint)

	id := c.Param("id")

	updatedExpense, err := h.uc.UpdateExpense(c.Request.Context(), input, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedExpense)
}

func (h *ExpenseHandler) UpdateExpenses(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var body struct {
		IDs    []string                   `json:"ids"`
		Values expense.UpdateExpenseInput `json:"values"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	body.Values.UserID = userID.(uint)

	updatedExpesnes, err := h.uc.UpdateExpenses(c.Request.Context(), body.Values, body.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedExpesnes)
}

func (h *ExpenseHandler) CreateExpenses(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var inputs []expense.CreateExpenseInput

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	for i := range inputs {
		inputs[i].UserID = userID.(uint)
	}

	createdExpenses, err := h.uc.CreateExpenses(c.Request.Context(), inputs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdExpenses})
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	userID, _ := c.Get("user_id")

	filters := parseFilters(c, userID.(uint))

	expenses, total, sum, err := h.uc.GetExpenses(c.Request.Context(), filters)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": expenses,
		"sum":  sum,
		"summary": gin.H{
			"total":     total,
			"page":      filters.Page,
			"page_size": filters.PageSize,
		},
	})
}

func (h *ExpenseHandler) GetExpensesByCategory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	filters := parseFilters(c, userID.(uint))

	groups, err := h.uc.GetExpensesByCategory(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

func (h *ExpenseHandler) GetExpensesByDate(c *gin.Context) {
	userID, _ := c.Get("user_id")
	filters := parseFilters(c, userID.(uint))

	groups, err := h.uc.GetExpensesByDate(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

func (h *ExpenseHandler) GetExpensesByDay(c *gin.Context) {
	userID, _ := c.Get("user_id")
	filters := parseFilters(c, userID.(uint))

	groups, err := h.uc.GetExpensesByDay(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

func (h *ExpenseHandler) GetExpensesByMonth(c *gin.Context) {
	userID, _ := c.Get("user_id")
	filters := parseFilters(c, userID.(uint))

	groups, err := h.uc.GetExpensesByMonth(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

func (h *ExpenseHandler) GetExpensesByYear(c *gin.Context) {
	userID, _ := c.Get("user_id")
	filters := parseFilters(c, userID.(uint))

	groups, err := h.uc.GetExpensesByYear(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	id := c.Param("id")

	err := h.uc.DeleteExpense(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *ExpenseHandler) DeleteExpenses(c *gin.Context) {
	var body struct {
		IDs []string `json:"ids"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	err := h.uc.DeleteExpenses(c.Request.Context(), body.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func parseFilters(c *gin.Context, userID uint) domain.ExpenseFilters {
	var filters domain.ExpenseFilters

	filters.UserID = userID

	if timestampStart := c.Query("timestamp_start"); timestampStart != "" {
		filters.TimestampStart = timestampStart
	}
	if timestampEnd := c.Query("timestamp_end"); timestampEnd != "" {
		filters.TimestampEnd = timestampEnd
	}
	if name := c.Query("name"); name != "" {
		filters.Name = name
	}
	if category := c.Query("category"); category != "" {
		filters.Category = category
	}

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "50")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	filters.Page = pageInt
	filters.PageSize = pageSizeInt

	if orderBy := c.Query("order_by"); orderBy != "" {
		filters.OrderBy = orderBy
	}

	if orderDirection := c.Query("order_direction"); orderDirection != "" {
		filters.OrderDirection = orderDirection
	}

	return filters
}
