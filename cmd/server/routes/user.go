package routes

import (
	"finance/config"
	"finance/internal/handler/http"
	authUC "finance/internal/usecase/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	respository "finance/internal/repository/user"
	usecase "finance/internal/usecase/user"
)

func UserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	userRepository := respository.NewPostgresRepository(db)
	userUseCase := usecase.NewUserUseCse(userRepository)
	authUseCase := authUC.NewAuthUseCase(cfg.JWTSecret)
	userHandler := http.NewUserHandler(userUseCase, authUseCase)

	router.GET("/user", userHandler.GetUsers)
	router.POST("/user", userHandler.CreateUser)
}
