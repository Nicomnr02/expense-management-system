package contextkey

type LocalKey string

const (
	User      LocalKey = "user"
	Worker    LocalKey = "worker"
	Logger    LocalKey = "logger"
	RequestID LocalKey = "request_id"
)
