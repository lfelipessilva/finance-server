package routes

import (
	"finance/config"
	"finance/internal/handler/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	usecase "finance/internal/usecase/auth"
)

func AuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	authUseCase := usecase.NewAuthUseCase(cfg.JWTSecret)
	authHandler := http.NewAuthHandler(authUseCase)

	router.POST("/auth", authHandler.Authenticate)
}
