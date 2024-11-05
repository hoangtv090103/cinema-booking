package authinterface

import (
	"bookingcinema/pkg/auth/authdomain"
	"context"
)

type IUserRepository interface {
	GetByID(ctx context.Context, id uint) (*authdomain.User, error)
	GetByEmail(ctx context.Context, email string) (*authdomain.User, error)
	GetRoles(ctx context.Context, id uint) ([]*authdomain.Role, error)
	GetPermissions(ctx context.Context, id uint) ([]*authdomain.Permission, error)
	GetAll(ctx context.Context) ([]*authdomain.User, error)
	Create(ctx context.Context, user *authdomain.UserCreate) error
	Update(ctx context.Context, user *authdomain.UserUpdate) error
	Delete(ctx context.Context, id uint) error
}
