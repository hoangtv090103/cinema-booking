package authinterface

import (
	"bookingcinema/pkg/auth/authdomain"
	"context"
)

// IPermissionRepository defines methods for accessing permissions
type IPermissionRepository interface {
	GetPermissionByID(ctx context.Context, id uint) (*authdomain.Permission, error)
	CreatePermission(ctx context.Context, permission *authdomain.PermissionCreate) error
	ListPermissions(ctx context.Context) ([]*authdomain.Permission, error)
}
