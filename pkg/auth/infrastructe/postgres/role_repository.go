package authinfra

import (
	"bookingcinema/pkg/auth/authdomain"
	authinterface "bookingcinema/pkg/auth/interfaces"
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) authinterface.IRoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) GetByID(ctx context.Context, id uint) (*authdomain.Role, error) {
	role := &authdomain.Role{}
	query := "SELECT id, name, created_at, updated_at FROM roles WHERE id = ?"
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	err := r.db.QueryRowContext(ctx, query, id).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	return role, err
}

func (r *RoleRepository) GetByName(ctx context.Context, name string) (*authdomain.Role, error) {
	role := &authdomain.Role{}
	query := "SELECT id, name, created_at, updated_at FROM roles WHERE name = ?"
	err := r.db.QueryRowContext(ctx, query, name).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	return role, err
}

func (r *RoleRepository) Create(ctx context.Context, role *authdomain.RoleCreate) error {
	query := "INSERT INTO roles (name) VALUES (?)"
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	return r.db.QueryRowContext(ctx, query, role.Name).Err()

}

func (r *RoleRepository) GetAll(ctx context.Context) ([]*authdomain.Role, error) {
	var roles []*authdomain.Role
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, created_at, updated_at FROM roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		role := &authdomain.Role{}
		if err := rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
