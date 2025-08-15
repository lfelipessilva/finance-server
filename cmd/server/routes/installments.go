package routes

import (
	"finance/internal/handler/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"finance/internal/repository/expense"
	repository "finance/internal/repository/installment"
	usecase "finance/internal/usecase/installment"
)

func InstallmentsRoutes(router *gin.RouterGroup, db *gorm.DB) {
	installmentRepository := repository.NewPostgresRepository(db)
	expenseRepository := expense.NewPostgresRepository(db)
	installmentUseCase := usecase.NewInstallmentUseCase(installmentRepository, expenseRepository)
	installmentHandler := http.NewInstallmentHandler(installmentUseCase)

	router.POST("/installments", installmentHandler.CreateInstallment)
}
