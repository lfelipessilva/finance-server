package routes

import (
	"finance/internal/handler/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	respository "finance/internal/repository/user"
	usecase "finance/internal/usecase/user"
)

func UserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := respository.NewPostgresRepository(db)
	userUseCase := usecase.NewUserUseCse(userRepository)
	userHandler := http.NewUserHandler(userUseCase)

	router.GET("/user", userHandler.GetUsers)
	router.POST("/user", userHandler.CreateUser)
}
