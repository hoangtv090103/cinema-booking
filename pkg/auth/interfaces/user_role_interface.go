package authinterface

import (
	"bookingcinema/pkg/auth/authdomain"
	"context"
)

type UserRoleRepository interface {
	AssignRoleToUser(ctx context.Context, userID uint, roleID uint) error
	GetUserRoles(ctx context.Context, userID uint) ([]*authdomain.Role, error)
	RemoveRoleFromUser(ctx context.Context, userID uint, roleID uint) error
}
