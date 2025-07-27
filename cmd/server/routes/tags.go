package routes

import (
	"finance/internal/handler/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repository "finance/internal/repository/tag"
	usecase "finance/internal/usecase/tag"
)

func TagsRoutes(router *gin.RouterGroup, db *gorm.DB) {
	tagRepository := repository.NewPostgresRepository(db)
	tagUseCase := usecase.NewTagUseCase(tagRepository)
	tagHandler := http.NewTagHandler(tagUseCase)

	router.GET("/tags", tagHandler.GetTags)
}
