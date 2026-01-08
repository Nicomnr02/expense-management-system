package authdomain

import "time"

type User struct {
	ID        int
	Email     string
	Name      string
	Role      string
	Password  string
	CreatedAt time.Time
}
