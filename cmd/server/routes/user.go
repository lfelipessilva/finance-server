package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	// cfg, err := config.Load()
	// if err != nil {
	// 	panic("failed to load config: " + err.Error())
	// }

	// userRepository := respository.NewPostgresRepository(db)
	// userUseCase := usecase.NewUserUseCse(userRepository)
	// authUseCase := authUC.NewAuthUseCase(cfg.JWTSecret, cfg.GoogleOAuthClientID, userUseCase)
	// userHandler := http.NewUserHandler(userUseCase, authUseCase)

	// router.GET("/user", userHandler.GetUsers)
	// router.POST("/user", userHandler.CreateUser)
}
