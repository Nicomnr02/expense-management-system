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
	AppPrefix       string
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

	SystemUserID      int
	MinExpenseAmount  int
	MaxExpenseAmount  int
	ApprovalThreshold int

	RedisAddr     string
	RedisMaxRetry int
	RedisTimeout  int

	PaymentURL string
}

func LoadConfig() *Config {
	if os.Getenv("APP_ENV") != "production" {
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

	systemUserID, _ := strconv.Atoi(getEnv("SYSTEM_USER_ID", "1"))
	minExpenseAmount, _ := strconv.Atoi(getEnv("MIN_EXPENSE_AMOUNT", "10000"))
	maxExpenseAmount, _ := strconv.Atoi(getEnv("MAX_EXPENSE_AMOUNT", "50000000"))
	approvalThreshold, _ := strconv.Atoi(getEnv("APPROVALTHRESHOLD", "1000000"))

	redisMaxRetry, _ := strconv.Atoi(getEnv("REDIS_MAX_RETRY", "3"))
	redisTimeout, _ := strconv.Atoi(getEnv("REDIS_TIMEOUT", "30"))

	return &Config{
		AppName:         getEnv("APP_NAME", "expenses-management-system"),
		AppPort:         getEnv("APP_PORT", "3000"),
		AppTimezone:     getEnv("APP_TIMEZONE", "Asia/Jakarta"),
		AppPrefix:       getEnv("APP_PREFIX", "/api"),
		AppReadTimeout:  AppReadTimeoutTime,
		AppWriteTimeout: AppWriteTimeoutTime,
		AppIdleTimeout:  AppIdleTimeoutTime,

		DBHost:              getEnv("DB_HOST", "localhost"),
		DBPort:              getEnv("DB_PORT", "5430"),
		DBName:              getEnv("DB_NAME", "expense_management"),
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

		SystemUserID:      systemUserID,
		MinExpenseAmount:  minExpenseAmount,
		MaxExpenseAmount:  maxExpenseAmount,
		ApprovalThreshold: approvalThreshold,

		RedisAddr:     getEnv("REDIS_ADDR", "redis:6379"),
		RedisMaxRetry: redisMaxRetry,
		RedisTimeout:  redisTimeout,

		PaymentURL: getEnv("PAYMENT_URL", "https://1620e98f-7759-431c-a2aa-f449d591150b.mock.pstmn.io/v1/payments"),
	}

}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
