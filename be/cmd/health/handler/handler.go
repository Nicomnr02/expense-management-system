package healthhandler

import (
	"context"
	"expense-management-system/cmd/health/enum"
	"expense-management-system/database"
	"expense-management-system/dto"
	"expense-management-system/pkg/httpserver"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type healthHandlerImpl struct {
	server   httpserver.Server
	database *database.Database
}

func New(server httpserver.Server, database *database.Database) {
	handler := healthHandlerImpl{
		server:   server,
		database: database,
	}

	server.App.Get("/health", server.Handler(handler.Check))
}

func (h *healthHandlerImpl) Check(ctx httpserver.Context) error {
	time := time.Now().Format(time.RFC3339)
	err := h.database.Conn.Ping(context.Background())
	if err != nil {
		return dto.Error(ctx, dto.ErrInternalServer(err.Error()), fiber.Map{
			"status": enum.DOWN,
			"time":   time,
			"deps": map[string]string{
				"database": enum.DOWN,
			},
		})
	}


	return dto.Success(ctx, http.StatusOK, fiber.Map{
		"status": enum.UP,
		"time":   time,
		"deps": map[string]string{
			"database": enum.UP,
		},
	})
}
