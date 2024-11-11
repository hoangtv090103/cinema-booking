package authdomain

import "time"

type User struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	RoleID      uint      `json:"role_id"`     // Foreign key to Role
	Permissions []string  `json:"permissions"` // Permissions granted to the user
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Active      bool      `json:"active"`
}

type UserCreate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleID   uint   `json:"-"`
	Password string `json:"password"`
}

type UserUpdate struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}
