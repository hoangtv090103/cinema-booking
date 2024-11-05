package authinfra

import (
	"bookingcinema/pkg/auth/authdomain"
	authinterface "bookingcinema/pkg/auth/interfaces"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PermissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) authinterface.IPermissionRepository {
	return &PermissionRepository{
		db: db,
	}
}

func (r *PermissionRepository) GetPermissionByID(ctx context.Context, permissionID uint) (*authdomain.Permission, error) {
	permission := &authdomain.Permission{}

	query := `SELECT id, create_perm, read_perm, update_perm, delete_perm FROM permissions WHERE id = ?`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	err := r.db.QueryRowContext(ctx, query, permissionID).Scan(
		&permission.ID,
		&permission.CreatePerm,
		&permission.ReadPerm,
		&permission.UpdatePerm,
		&permission.DeletePerm,
	)

	if err != nil {
		return nil, fmt.Errorf("Cannot get permissions: %v", err)
	}

	return permission, nil
}

func (r *PermissionRepository) CreatePermission(ctx context.Context, permission *authdomain.PermissionCreate) error {
	query := `INSERT INTO permissions(create_perm, read_perm, update_perm, delete_perm) VALUES (?, ?, ?, ?)`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	err := r.db.QueryRowContext(
		ctx,
		query,
		permission.CreatePerm,
		permission.ReadPerm,
		permission.UpdatePerm,
		permission.DeletePerm,
	).Err()
	return err
}

func (r *PermissionRepository) ListPermissions(ctx context.Context) ([]*authdomain.Permission, error) {
	var permissions []*authdomain.Permission
	query := `SELECT id, create_perm, read_perm, update_perm, delete_perm, created_at, updated_at, active FROM permissions WHERE active = true`

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot list all permissions: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		per := &authdomain.Permission{}
		err := rows.Scan(&per.ID, &per.CreatePerm, &per.ReadPerm, &per.UpdatePerm, &per.DeletePerm, &per.CreatedAt, &per.UpdatedAt, &per.Active)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, per)
	}

	return permissions, nil
}
