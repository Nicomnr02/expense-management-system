package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName         string
	AppEnv          string
	AppPort         string
	AppTimezone     string
	AppReadTimeout  time.Duration
	AppWriteTimeout time.Duration
	AppIdleTimeout  time.Duration

	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBMaxConns          int
	DBMaxConnIdle       int
	DBConnectionTimeout int
	DBSSLMode           string

	JWTAccessSecret                string
	JWTRefreshSecret               string
	JWTAccessTokenDurationMinutes  time.Duration
	JWTRefreshTokenDurationMinutes time.Duration
}

func LoadConfig() *Config {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("No .env file found (using system envs)")
		}
	}
	
	AppReadTimeout, err := strconv.Atoi(getEnv("APP_READ_TIMEOUT", "10"))
	if err != nil {
		log.Fatal("invalid APP_READ_TIMEOUT:", err)
	}
	AppWriteTimeout, err := strconv.Atoi(getEnv("APP_WRITE_TIMEOUT", "10"))
	if err != nil {
		log.Fatal("invalid APP_WRITE_TIMEOUT:", err)
	}
	AppIdleTimeout, err := strconv.Atoi(getEnv("APP_IDLE_TIMEOUT", "60"))
	if err != nil {
		log.Fatal("invalid APP_IDLE_TIMEOUT:", err)
	}
	AppReadTimeoutTime := time.Duration(AppReadTimeout) * time.Minute
	AppWriteTimeoutTime := time.Duration(AppWriteTimeout) * time.Minute
	AppIdleTimeoutTime := time.Duration(AppIdleTimeout) * time.Minute

	DBMaxConns, _ := strconv.Atoi(getEnv("DB_MAX_CONNS", "10"))
	DBMaxConnIdle, _ := strconv.Atoi(getEnv("DB_MAX_CONNS_IDLE", "5"))
	DBConnectionTimeout, _ := strconv.Atoi(getEnv("DB_CONNECTION_TIMEOUT", "5"))

	jwtAccessTokenDurationStr := getEnv("JWT_ACCESS_TOKEN_DURATION_MINUTES", "5")
	jwtRefreshTokenDurationStr := getEnv("JWT_REFRESH_TOKEN_DURATION_MINUTES", "60")
	jwtAccessTokenDuration, err := strconv.Atoi(jwtAccessTokenDurationStr)
	if err != nil {
		log.Fatal("invalid JWT_ACCESS_TOKEN_DURATION_MINUTES:", err)
	}
	jwtRefreshTokenDuration, err := strconv.Atoi(jwtRefreshTokenDurationStr)
	if err != nil {
		log.Fatal("invalid JWT_REFRESH_TOKEN_DURATION_MINUTES:", err)
	}
	jwtAccessTokenDurationTime := time.Duration(jwtAccessTokenDuration) * time.Minute
	jwtRefreshTokenDurationTime := time.Duration(jwtRefreshTokenDuration) * time.Minute

	return &Config{
		AppName:         getEnv("APP_NAME", "expenses-management-system"),
		AppPort:         getEnv("APP_PORT", "3000"),
		AppTimezone:     getEnv("APP_TIMEZONE", "Asia/Jakarta"),
		AppReadTimeout:  AppReadTimeoutTime,
		AppWriteTimeout: AppWriteTimeoutTime,
		AppIdleTimeout:  AppIdleTimeoutTime,

		DBHost:              getEnv("DB_HOST", "localhost"),
		DBPort:              getEnv("DB_PORT", "5430"),
		DBUser:              getEnv("DB_USER", "postgres"),
		DBPassword:          getEnv("DB_PASSWORD", "sec"),
		DBSSLMode:           getEnv("DB_SSL_MODE", "disable"),
		DBMaxConns:          DBMaxConns,
		DBMaxConnIdle:       DBMaxConnIdle,
		DBConnectionTimeout: DBConnectionTimeout,

		JWTAccessSecret:                getEnv("JWT_ACCESS_SECRET", "supersecret"),
		JWTRefreshSecret:               getEnv("JWT_REFRESH_SECRET", "supersecret"),
		JWTAccessTokenDurationMinutes:  jwtAccessTokenDurationTime,
		JWTRefreshTokenDurationMinutes: jwtRefreshTokenDurationTime,
	}

}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
