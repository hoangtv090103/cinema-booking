package authusecase

import (
	"bookingcinema/pkg/auth/authdomain"
	authinterface "bookingcinema/pkg/auth/interfaces"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type IAuthenticationUseCase interface {
	Register(ctx context.Context, user *authdomain.UserCreate) error
	Login(ctx context.Context, email, password string) (*authdomain.User, error)
}

type AuthenticationUseCase struct {
	userRepo authinterface.IUserRepository
}

func NewAuthenticationUseCase(userRepo authinterface.IUserRepository) IAuthenticationUseCase {
	return &AuthenticationUseCase{
		userRepo: userRepo,
	}
}

func (u *AuthenticationUseCase) Login(ctx context.Context, email, password string) (*authdomain.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("cannot find the user with %s: %v", email, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (u *AuthenticationUseCase) Register(ctx context.Context, user *authdomain.UserCreate) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("cannot hash password: %v", err)
	}

	user.Password = string(hashedPassword)
	user.RoleID = 2
	return u.userRepo.Create(ctx, user)
}
