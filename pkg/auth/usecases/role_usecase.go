package authusecase

import (
	"bookingcinema/pkg/auth/authdomain"
	authinterface "bookingcinema/pkg/auth/interfaces"
	"context"
)

type IRoleUseCase interface {
	GetByName(ctx context.Context, name string) (*authdomain.Role, error)
}

type RoleUseCase struct {
	roleRepo authinterface.IRoleRepository
}

func (u *RoleUseCase) GetRoleByName(ctx context.Context, name string) (*authdomain.Role, error) {
	return u.roleRepo.GetByName(ctx, name)
}
