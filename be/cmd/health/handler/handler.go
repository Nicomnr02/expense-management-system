package healthhandler

import (
	"context"
	"expense-management-system/cmd/health/enum"
	"expense-management-system/database"
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type healthHandlerImpl struct {
	router   fiber.Router
	database *database.Database
}

func New(server *fiber.App, database *database.Database) {
	handler := healthHandlerImpl{
		router:   server,
		database: database,
	}

	server.Get("/health", handler.Check)
}

func (h *healthHandlerImpl) Check(c *fiber.Ctx) error {
	log := c.Locals(contextkey.Logger).(*zap.Logger)
	time := time.Now().Format(time.RFC3339)
	err := h.database.Conn.Ping(context.Background())
	if err != nil {
		log.Error(err.Error())
		return model.Error(c, model.ErrInternalServer(err.Error()), fiber.Map{
			"status": enum.DOWN,
			"time":   time,
			"deps": map[string]string{
				"database": enum.DOWN,
			},
		})
	}

	return model.Success(c, http.StatusOK, fiber.Map{
		"status": enum.UP,
		"time":   time,
		"deps": map[string]string{
			"database": enum.UP,
		},
	})
}
