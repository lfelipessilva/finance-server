package main

import (
	"finance/cmd/server/routes"
	"finance/config"
	"finance/internal/middleware"
	userRepo "finance/internal/repository/user"
	userUC "finance/internal/usecase/user"

	"finance/pkg/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode)

	// Create database config with pool settings
	dbConfig := database.DatabaseConfig{
		DSN:             dsn,
		MaxOpenConns:    cfg.DBMaxOpenConns,
		MaxIdleConns:    cfg.DBMaxIdleConns,
		ConnMaxLifetime: cfg.DBConnMaxLifetime,
		ConnMaxIdleTime: cfg.DBConnMaxIdleTime,
	}

	db, err := database.NewPostgresConnectionWithConfig(&dbConfig)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

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

	// Health check endpoint with database pool stats
	router.GET("/health", func(c *gin.Context) {
		stats, err := database.GetConnectionStats(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"database": gin.H{
				"pool_stats": stats,
			},
		})
	})

	api := router.Group("/api/v1")
	{
		// Public routes (no auth required)
		routes.AuthRoutes(api, db)
		routes.CategoriesRoutes(api, db)
		routes.TagsRoutes(api, db)

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
