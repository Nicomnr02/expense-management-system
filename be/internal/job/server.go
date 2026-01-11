package job

import (
	"context"
	"expense-management-system/config"
	"expense-management-system/internal/contextkey"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type Server interface {
	Ping() error
	Run() error
	Shutdown()
	RegisterWorker(action string, f func(c context.Context, task Task) error)
}
type server struct {
	server *asynq.Server
	mux    *asynq.ServeMux
	log    *zap.Logger
}

func (s *server) Ping() error {
	return s.server.Ping()
}

func (s *server) Run() error {
	return s.server.Run(s.mux)
}

func (s *server) Shutdown() {
	s.server.Shutdown()
}

func NewServer(cfg *config.Config, log *zap.Logger) Server {
	conn := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: cfg.RedisAddr,
		},
		asynq.Config{
			Concurrency: 5,
		},
	)

	return &server{
		server: conn,
		mux:    asynq.NewServeMux(),
		log:    log,
	}
}

type Worker func(c context.Context, t *asynq.Task) error

func (s *server) RegisterWorker(action string, f func(c context.Context, task Task) error) {
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
