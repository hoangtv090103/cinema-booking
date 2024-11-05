package authinfra

import (
	"bookingcinema/pkg/auth/authdomain"
	authinterface "bookingcinema/pkg/auth/interfaces"
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type RolePermissionRepository struct {
	db *sql.DB
}

func NewRolePermissionRepository(db *sql.DB) authinterface.IRolePermissionRepository {
	return &RolePermissionRepository{
		db: db,
	}
}

// AssignPermissionToRole assigns a permission to a role by inserting into role_permissions table
func (r *RolePermissionRepository) AssignPermissionToRole(ctx context.Context, roleID uint, permissionID uint) error {
	query := "INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?) ON CONFLICT DO NOTHING"
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err := r.db.ExecContext(ctx, query, roleID, permissionID)
	return err
}

// GetRolePermissions retrieves all permissions assigned to a role
func (r *RolePermissionRepository) GetRolePermissions(ctx context.Context, roleID uint) ([]authdomain.Permission, error) {
	var permissions []authdomain.Permission
	query := `
        SELECT p.id, p.create_perm, p.read_perm, p.update_perm, p.delete_perm, p.created_at, p.updated_at
        FROM permissions p
        INNER JOIN role_permissions rp ON rp.permission_id = p.id
        WHERE rp.role_id = ?
    `
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err := r.db.QueryContext(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission authdomain.Permission
		if err := rows.Scan(&permission.ID, &permission.CreatePerm, &permission.ReadPerm, &permission.UpdatePerm, &permission.DeletePerm, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

// RemovePermissionFromRole removes a specific permission from a role
func (r *RolePermissionRepository) RemovePermissionFromRole(ctx context.Context, roleID uint, permissionID uint) error {
	query := "DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?"
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err := r.db.ExecContext(ctx, query, roleID, permissionID)
	return err
}
