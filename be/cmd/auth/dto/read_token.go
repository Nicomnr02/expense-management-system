package authdto

type ReadTokenRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
