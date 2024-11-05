package authdomain

import "time"

type Permission struct {
	ID         uint      `json:"id"`
	CreatePerm bool      `json:"create_perm"`
	ReadPerm   bool      `json:"read_perm"`
	UpdatePerm bool      `json:"update_perm"`
	DeletePerm bool      `json:"delete_perm"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Active     bool      `json:"active"`
}

type PermissionCreate struct {
	CreatePerm bool `json:"create_perm"`
	ReadPerm   bool `json:"read_perm"`
	UpdatePerm bool `json:"update_perm"`
	DeletePerm bool `json:"delete_perm"`
}
