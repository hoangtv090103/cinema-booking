package authinfra

import (
	"bookingcinema/pkg/auth/authdomain"
	authinterface "bookingcinema/pkg/auth/interfaces"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRoleRepository struct {
	db *sql.DB
}

func NewUserRoleRepository(db *sql.DB) authinterface.IUserRoleRepository {
	return &UserRoleRepository{
		db: db,
	}
}

func (r *UserRoleRepository) AssignRoleToUser(ctx context.Context, userID uint, roleID uint) error {
	query := `INSERT INTO user_roles (user_id, role_id) VALUES (?, ?) ON CONFLICT DO NOTHING`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	_, err := r.db.ExecContext(ctx, query, userID, roleID)

	return fmt.Errorf("cannot assign role to user: %v", err)
}

func (r *UserRoleRepository) GetUserRoles(ctx context.Context, userID uint) ([]*authdomain.Role, error) {
	var roles []*authdomain.Role

	query := `
			SELECT r.id, r.name, r.created_at, r.updated_at 
			FROM roles r 
			INNER JOIN user_roles ur ON ur.role_id = r.id 
			WHERE ur.user_id = ?
			`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("cannot get user's roles: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		role := &authdomain.Role{}
		err = rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}
