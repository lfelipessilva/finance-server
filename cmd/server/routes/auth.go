package routes

import (
	"finance/config"
	"finance/internal/handler/http"
	userRepo "finance/internal/repository/user"
	userUC "finance/internal/usecase/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	usecase "finance/internal/usecase/auth"
)

func AuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	userRepository := userRepo.NewPostgresRepository(db)
	userUseCase := userUC.NewUserUseCse(userRepository)
	authUseCase := usecase.NewAuthUseCase(cfg.JWTSecret, cfg.GoogleOAuthClientID, userUseCase)
	authHandler := http.NewAuthHandler(authUseCase)

	router.POST("/auth", authHandler.Authenticate)
}
