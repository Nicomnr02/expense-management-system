package job

type Task struct {
	Action     string `json:"action"`
	Payload    []byte `json:"payload"`
	RetryCount int    `json:"retry_count"`
	MaxRetry   int    `json:"max_retry"`
}
