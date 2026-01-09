package httpserver

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Context *fiber.Ctx

func UseContext(c Context) context.Context {
	var ctx context.Context
	_ = func(c *fiber.Ctx) error {
		ctx = c.Context()
		return nil
	}(c)

	return ctx
}

func UseLogger(c Context) *zap.Logger {
	var logger *zap.Logger
	_ = func(c *fiber.Ctx) error {
		_logger, ok := c.Locals("logger").(*zap.Logger)
		if ok {
			_logger.WithOptions(zap.AddCaller())
			logger = _logger
		}
		return nil
	}(c)

	return logger
}
