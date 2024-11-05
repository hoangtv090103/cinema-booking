package authinterface

import (
	"bookingcinema/pkg/auth/authdomain"
	"context"
)

// IUserRoleRepository defines methods for managing user-role associations
type IUserRoleRepository interface {
	AssignRoleToUser(ctx context.Context, userID uint, roleID uint) error
	GetUserRoles(ctx context.Context, userID uint) ([]*authdomain.Role, error)
}
