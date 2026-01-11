package job

import (
	"expense-management-system/config"
	"time"

	"github.com/hibiken/asynq"
)

type Client struct {
	conn *asynq.Client
	cfg  *config.Config
}

func NewClient(cfg *config.Config) *Client {
	conn := asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.RedisAddr,
	})

	return &Client{conn, cfg}
}

func (c *Client) Enqueue(t Task) error {
	task := asynq.NewTask(
		t.Action,
		t.Payload,
		asynq.MaxRetry(c.cfg.RedisMaxRetry),
		asynq.Timeout(time.Duration(c.cfg.RedisTimeout)*time.Second),
	)

	_, err := c.conn.Enqueue(task)

	return err
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Ping() error {
	return c.conn.Ping()
}
