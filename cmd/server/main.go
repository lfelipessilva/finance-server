package main

import (
	"finance/config"
	"finance/internal/handler/http"
	expenseRepo "finance/internal/repository/expense"
	expenseUseCase "finance/internal/usecase/expense"
	"finance/pkg/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	// Database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode)
	db, err := database.NewPostgresConnection(dsn)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Run migrations (you might want to run this separately)
	// migrate -path migrations -database "postgres://user:pass@host:port/dbname?sslmode=disable" up

	// Setup layers
	expenseRepo := expenseRepo.NewPostgresRepository(db)
	expenseUC := expenseUseCase.NewExpenseUseCase(expenseRepo)
	expenseHandler := http.NewExpenseHandler(expenseUC)

	// Server setup
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	api := router.Group("/api/v1")
	{
		api.POST("/expenses", expenseHandler.CreateExpense)
		api.PUT("/expenses/:id", expenseHandler.UpdateExpense)
		api.POST("/expenses/batch", expenseHandler.CreateExpenses)
		api.GET("/expenses", expenseHandler.GetExpenses)
	}

	if err := router.Run(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
