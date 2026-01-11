package mocks

import (
	"expense-management-system/config"
	"time"
)

func Config() *config.Config {
	return &config.Config{
		AppName:         "test-app",
		AppPort:         "3001",
		AppTimezone:     "UTC",
		AppReadTimeout:  time.Second * 10,
		AppWriteTimeout: time.Second * 10,
		AppIdleTimeout:  time.Second * 30,

		DBHost:              "localhost",
		DBPort:              "5432",
		DBUser:              "postgres",
		DBPassword:          "secret",
		DBName:              "test_db",
		DBSSLMode:           "disable",
		DBMaxConns:          5,
		DBMaxConnIdle:       2,
		DBConnectionTimeout: 5,

		JWTAccessSecret:                "test-secret",
		JWTRefreshSecret:               "test-secret",
		JWTAccessTokenDurationMinutes:  time.Minute * 5,
		JWTRefreshTokenDurationMinutes: time.Minute * 60,

		SystemUserID:      1,
		MinExpenseAmount:  10000,
		MaxExpenseAmount:  5000000,
		ApprovalThreshold: 1000000,

		RedisAddr:     "localhost:6379",
		RedisMaxRetry: 3,
		RedisTimeout:  30,

		PaymentURL: "http://mock-payment.test",
	}
}
