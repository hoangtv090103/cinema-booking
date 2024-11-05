package authinterface

import (
	"bookingcinema/pkg/auth/authdomain"
	"context"
)

// IRolePermissionRepository defines methods for managing role-permission associations
type IRolePermissionRepository interface {
	AssignPermissionToRole(ctx context.Context, roleID uint, permissionID uint) error
	GetRolePermissions(ctx context.Context, roleID uint) ([]authdomain.Permission, error)
}
