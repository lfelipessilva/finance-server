package main

import (
	"finance/cmd/server/routes"
	"finance/config"
	"finance/internal/handler/http"
	"finance/internal/middleware"
	catRepo "finance/internal/repository/category"
	tagRepo "finance/internal/repository/tag"
	userRepo "finance/internal/repository/user"
	catUC "finance/internal/usecase/category"
	tagUC "finance/internal/usecase/tag"
	userUC "finance/internal/usecase/user"

	"finance/pkg/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode)
	db, err := database.NewPostgresConnection(dsn)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	tagRepository := tagRepo.NewPostgresRepository(db)
	tagUseCase := tagUC.NewTagUseCase(tagRepository)
	tagHandler := http.NewTagHandler(tagUseCase)

	categoryRepository := catRepo.NewPostgresRepository(db)
	categoryUseCase := catUC.NewCategoryUseCse(categoryRepository)
	categoryHandler := http.NewCategoryHandler(categoryUseCase)

	userRepository := userRepo.NewPostgresRepository(db)
	userUseCase := userUC.NewUserUseCse(userRepository)
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret, userUseCase)

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
		// Public routes (no auth required)
		routes.AuthRoutes(api, db)
		api.GET("/categories", categoryHandler.GetCategories)
		api.GET("/tags", tagHandler.GetTags)

		// Protected routes (auth required)
		protected := api.Group("")
		protected.Use(authMiddleware.Authenticate())
		{
			routes.ExpensesRoutes(protected, db)
			routes.UserRoutes(protected, db)
		}
	}

	if err := router.Run(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
