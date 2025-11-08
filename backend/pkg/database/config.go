package database

import (
	"fmt"
	"os"
)

// DBConfig holds database configuration
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// GetDBConfig returns database configuration from environment variables
func GetDBConfig() DBConfig {
	config := DBConfig{
		User:     getEnvOrDefault("DB_USER", "root"),
		Password: getEnvOrDefault("DB_PASSWORD", ""),
		Host:     getEnvOrDefault("DB_HOST", "127.0.0.1"),
		Port:     getEnvOrDefault("DB_PORT", "3306"),
		Name:     getEnvOrDefault("DB_NAME", "spk_profile_matching"),
	}
	return config
}

// GetTestDBConfig returns test database configuration
func GetTestDBConfig() DBConfig {
	config := GetDBConfig()

	// Override database name for test
	if testDBName := os.Getenv("TEST_DB_NAME"); testDBName != "" {
		config.Name = testDBName
	} else {
		config.Name = config.Name + "_test"
	}

	return config
}

// BuildDSN builds MySQL DSN string from config
func (c DBConfig) BuildDSN() string {
	if c.Password != "" {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.User, c.Password, c.Host, c.Port, c.Name)
	}
	return fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Host, c.Port, c.Name)
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
