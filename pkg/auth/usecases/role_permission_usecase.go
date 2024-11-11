package authusecase

import (
	"bookingcinema/pkg/auth/authdomain"
	authinterface "bookingcinema/pkg/auth/interfaces"
	"context"
)

type IRolePermissionUseCase interface {
	AssignRoleToUser(ctx context.Context, userID uint, roleID uint) error
	AssignPermissionToRole(ctx context.Context, roleID, permissionID uint) error
	GetRolePermissions(ctx context.Context, roleID uint) ([]*authdomain.Permission, error)
}

type RolePermissionUseCase struct {
	roleRepo           authinterface.IPermissionRepository
	permissionRepo     authinterface.IPermissionRepository
	userRoleRepo       authinterface.IUserRoleRepository
	rolePermissionRepo IRolePermissionUseCase
}

func NewRolePermissionUseCase(
	roleRepo authinterface.IPermissionRepository,
	permissionRepo authinterface.IPermissionRepository,
	userRoleRepo authinterface.IUserRoleRepository,
	rolePermissionRepo IRolePermissionUseCase,
) IRolePermissionUseCase {
	return &RolePermissionUseCase{
		roleRepo:           roleRepo,
		permissionRepo:     permissionRepo,
		userRoleRepo:       userRoleRepo,
		rolePermissionRepo: rolePermissionRepo,
	}
}

// AssignRoleToUser assigns a role to a user
func (uc *RolePermissionUseCase) AssignRoleToUser(ctx context.Context, userID uint, roleID uint) error {
	return uc.userRoleRepo.AssignRoleToUser(ctx, userID, roleID)
}

// AssignPermissionToRole assigns a permission to a role
func (uc *RolePermissionUseCase) AssignPermissionToRole(ctx context.Context, roleID, permissionID uint) error {
	return uc.rolePermissionRepo.AssignPermissionToRole(ctx, roleID, permissionID)
}

// GetUserRoles retrieves all roles assigned to a user
func (uc *RolePermissionUseCase) GetUserRoles(ctx context.Context, userID uint) ([]*authdomain.Role, error) {
	return uc.userRoleRepo.GetUserRoles(ctx, userID)
}

// GetRolePermissions retrieves all permissions assigned to a role
func (uc *RolePermissionUseCase) GetRolePermissions(ctx context.Context, roleID uint) ([]*authdomain.Permission, error) {
	return uc.rolePermissionRepo.GetRolePermissions(ctx, roleID)
}
