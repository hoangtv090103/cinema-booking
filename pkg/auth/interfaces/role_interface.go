package authinterface

import (
	"bookingcinema/pkg/auth/authdomain"
	"context"
)

type IRoleRepository interface {
	GetByID(ctx context.Context, id uint) (*authdomain.Role, error)
	GetByName(ctx context.Context, name string) (*authdomain.Role, error)

	Create(ctx context.Context, role *authdomain.RoleCreate) error
	GetAll(ctx context.Context) ([]*authdomain.Role, error)
}
