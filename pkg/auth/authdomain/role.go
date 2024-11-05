package authdomain

import "time"

type Role struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool      `json:"active"`
}

type RoleCreate struct {
	Name string `json:"name"`
}
