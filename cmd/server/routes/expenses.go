package routes

import (
	"finance/internal/handler/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repository "finance/internal/repository/expense"
	tRepository "finance/internal/repository/tag"
	usecase "finance/internal/usecase/expense"
)

func ExpensesRoutes(router *gin.RouterGroup, db *gorm.DB) {
	expenseRepository := repository.NewPostgresRepository(db)
	tagRepository := tRepository.NewPostgresRepository(db)
	expenseUseCase := usecase.NewExpenseUseCase(expenseRepository, tagRepository)
	expenseHandler := http.NewExpenseHandler(expenseUseCase)

	router.GET("/expenses", expenseHandler.GetExpenses)
	router.GET("/expenses/category", expenseHandler.GetExpensesByCategory)
	router.GET("/expenses/date", expenseHandler.GetExpensesByDate)
	router.GET("/expenses/day", expenseHandler.GetExpensesByDay)
	router.GET("/expenses/month", expenseHandler.GetExpensesByMonth)
	router.GET("/expenses/year", expenseHandler.GetExpensesByYear)
	router.POST("/expenses/batch", expenseHandler.CreateExpenses)
	router.POST("/expenses", expenseHandler.CreateExpense)
	router.PUT("/expenses/batch", expenseHandler.UpdateExpenses)
	router.PUT("/expenses/:id", expenseHandler.UpdateExpense)
	router.DELETE("/expenses/batch", expenseHandler.DeleteExpenses)
	router.DELETE("/expenses/:id", expenseHandler.DeleteExpense)
}
