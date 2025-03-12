package http

import (
	"finance/internal/domain/vo"
	"finance/internal/usecase/expense"
	"net/http"
	"strconv"
	"time"

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

	expenses, err := h.uc.GetExpenses(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func parseFilters(c *gin.Context) expense.ExpenseFilters {
	var filters expense.ExpenseFilters

	if month := c.Query("month"); month != "" {
		year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		monthInt, _ := strconv.Atoi(month)
		yearInt, _ := strconv.Atoi(year)

		if my, err := vo.NewMonthYear(monthInt, yearInt); err == nil {
			filters.MonthYear = &my
		}
	}

	filters.Category = c.Query("category")
	return filters
}
