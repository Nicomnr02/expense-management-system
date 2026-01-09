package contextkey

type LocalKey string

const (
	User      LocalKey = "user"
	Logger    LocalKey = "logger"
	RequestID LocalKey = "request_id"
)
