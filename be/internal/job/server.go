package job

import (
	"context"
	"expense-management-system/config"
	"expense-management-system/internal/contextkey"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type Server struct {
	server *asynq.Server
	mux    *asynq.ServeMux
	log    *zap.Logger
}

func (s *Server) Ping() error {
	return s.server.Ping()
}

func (s *Server) Run() error {
	return s.server.Run(s.mux)
}

func (s *Server) Shutdown() {
	s.server.Shutdown()
}

func NewServer(cfg *config.Config, log *zap.Logger) *Server {
	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: cfg.RedisAddr,
		},
		asynq.Config{
			Concurrency: 5,
		},
	)

	return &Server{
		server: server,
		mux:    asynq.NewServeMux(),
		log:    log,
	}
}

type Worker func(c context.Context, t *asynq.Task) error

func (s *Server) RegisterWorker(action string, f func(c context.Context, task Task) error) {
	s.mux.HandleFunc(action, func(c context.Context, t *asynq.Task) error {
		task := Task{
			Payload: t.Payload(),
		}

		retried, ok := asynq.GetRetryCount(c)
		if ok {
			task.RetryCount = retried
		}

		maxRetry, ok := asynq.GetMaxRetry(c)
		if ok {
			task.MaxRetry = maxRetry
		}

		taskLogger := s.log.With(zap.String("task_id", t.ResultWriter().TaskID()))

		c = context.WithValue(c, contextkey.Worker, taskLogger)

		return f(c, task)
	})
}
