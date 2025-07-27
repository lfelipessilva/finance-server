package routes

import (
	"finance/internal/handler/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repository "finance/internal/repository/category"
	usecase "finance/internal/usecase/category"
)

func CategoriesRoutes(router *gin.RouterGroup, db *gorm.DB) {
	categoryRepository := repository.NewPostgresRepository(db)
	categoryUseCase := usecase.NewCategoryUseCse(categoryRepository)
	categoryHandler := http.NewCategoryHandler(categoryUseCase)

	router.GET("/categories", categoryHandler.GetCategories)
}
