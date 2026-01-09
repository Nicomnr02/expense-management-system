package httpserver

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

type ServerConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func (s *Server) Shutdown() {
	_ = s.App.Shutdown()
}

func (s *Server) Group(domain string) fiber.Router {
	return s.App.Group(domain)
}

func (s *Server) Use(fn func(Context) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return fn(c)
	}
}

func (s *Server) Middleware(m fiber.Handler) {
	_ = s.App.Use(m)
}

func (s *Server) Parse(c Context, out any) error {
	return func(c *fiber.Ctx) error {
		return c.BodyParser(&out)
	}(c)
}

func New(cfg ServerConfig) *Server {
	app := fiber.New(
		fiber.Config{
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	)
	return &Server{
		App: app,
	}
}
