package healthdto

type HealthRes struct {
	Status string        `json:"status"`
	Time   string        `json:"time"`
	Deps   HealthDepsRes `json:"deps"`
}

type HealthDepsRes struct {
	Database string `json:"database"`
	Queue    string `json:"queue"`
}
