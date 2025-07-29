package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	DSN string
	// Pool settings
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	// Logger settings
	LogLevel logger.LogLevel
}

// DefaultDatabaseConfig returns a default database configuration
func DefaultDatabaseConfig(dsn string) *DatabaseConfig {
	return &DatabaseConfig{
		DSN:             dsn,
		MaxOpenConns:    25,              // Maximum number of open connections
		MaxIdleConns:    10,              // Maximum number of idle connections
		ConnMaxLifetime: 5 * time.Minute, // Maximum lifetime of a connection
		ConnMaxIdleTime: 1 * time.Minute, // Maximum idle time of a connection
		LogLevel:        logger.Warn,     // Only log warnings and errors
	}
}

func NewPostgresConnection(dsn string) (*gorm.DB, error) {
	return NewPostgresConnectionWithConfig(DefaultDatabaseConfig(dsn))
}

func NewPostgresConnectionWithConfig(config *DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogLevel),
	})
	if err != nil {
		return nil, err
	}

	// Get the underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	return db, nil
}

// GetConnectionStats returns current connection pool statistics
func GetConnectionStats(db *gorm.DB) (map[string]interface{}, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	stats := sqlDB.Stats()
	return map[string]interface{}{
		"max_open_connections": stats.MaxOpenConnections,
		"open_connections":     stats.OpenConnections,
		"in_use":               stats.InUse,
		"idle":                 stats.Idle,
		"wait_count":           stats.WaitCount,
		"wait_duration":        stats.WaitDuration,
		"max_idle_closed":      stats.MaxIdleClosed,
		"max_lifetime_closed":  stats.MaxLifetimeClosed,
	}, nil
}
