package healthhandler

import (
	"context"
	healthdto "expense-management-system/cmd/health/dto"
	"expense-management-system/cmd/health/enum"
	"expense-management-system/database"
	"expense-management-system/internal/contextkey"
	"expense-management-system/internal/job"
	"expense-management-system/model"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type healthHandlerImpl struct {
	router    fiber.Router
	database  *database.Database
	jobClient *job.Client
	jobServer *job.Server
}

func New(
	server *fiber.App,
	database *database.Database,
	jobServer *job.Server,
) {
	handler := healthHandlerImpl{
		router:    server,
		database:  database,
		jobServer: jobServer,
	}

	server.Get("/health", handler.Check)
}

func (h *healthHandlerImpl) Check(c *fiber.Ctx) error {
	log := c.Locals(contextkey.Logger).(*zap.Logger)
	time := time.Now().Format(time.RFC3339)
	ctx := context.Background()
	data := healthdto.HealthRes{
		Status: enum.UP,
		Time:   time,
		Deps: healthdto.HealthDepsRes{
			Database: enum.UP,
			Queue:    enum.UP,
		},
	}

	if err := h.database.Conn.Ping(ctx); err != nil {
		log.Error(err.Error())
		data.Status = enum.DOWN
		data.Deps.Database = enum.DOWN

	} else if err := h.jobServer.Ping(); err != nil {
		log.Error(err.Error())
		data.Status = enum.DEGRADED
		data.Deps.Queue = enum.DOWN
	}

	return model.Success(c, http.StatusOK, data)
}
