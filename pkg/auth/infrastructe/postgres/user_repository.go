package authinfra

import (
	"bookingcinema/pkg/auth/authdomain"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*authdomain.User, error) {
	user := &authdomain.User{}

	query := `SELECT id, name, email, created_at, updated_at, active FROM users WHERE id=?`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Active,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*authdomain.User, error) {
	user := &authdomain.User{}
	query := `SELECT id, name, email, password, created_at, updated_at, active FROM users WHERE email=?`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot get user by email: %v", err)
	}

	return user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*authdomain.User, error) {
	var users []*authdomain.User

	query := `SELECT id, name, email, created_at, updated_at, active FROM users`

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &authdomain.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Active); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Create(ctx context.Context, user *authdomain.UserCreate) error {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?) returning id`
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	var userID uint
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&userID)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *authdomain.UserUpdate) error {
	query := `UPDATE users SET `
	args := []interface{}{}

	if user.Name != "" {
		query += "name = ?, "
		args = append(args, user.Name)
	}
	if user.Email != "" {
		query += "email = ?, "
		args = append(args, user.Email)
	}
	if user.Password != "" {
		query += "password = ?, "
		args = append(args, user.Password)
	}
	if user.Active {
		query += "active = ?, "
		args = append(args, user.Active)
	}

	// Remove the trailing comma and space
	query = query[:len(query)-2]
	query += " WHERE id = ?"
	args = append(args, user.ID)

	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	query := `UPDATE users SET active = false WHERE id = ?`
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepository) GetRoles(ctx context.Context, userID uint) ([]*authdomain.Role, error) {
	var roles []*authdomain.Role

	query := `
        SELECT r.id, r.name, r.created_at, r.updated_at
        FROM roles r
        INNER JOIN user_roles ur ON ur.role_id = r.id
        WHERE ur.user_id = ?
    `
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	rows, err := r.db.QueryContext(ctx, query, userID)

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

// GetUserPermissions retrieves the permissions associated with the user's roles.
func (r *UserRepository) GetPermissions(ctx context.Context, userID uint) ([]*authdomain.Permission, error) {
	var permissions []*authdomain.Permission

	query := `
        SELECT DISTINCT p.name
        FROM permissions p
        INNER JOIN role_permissions rp ON rp.permission_id = p.id
        INNER JOIN user_roles ur ON ur.role_id = rp.role_id
        WHERE ur.user_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		permission := &authdomain.Permission{}
		if err := rows.Scan(&permission); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
