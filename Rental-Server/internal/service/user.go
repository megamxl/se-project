package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
)

type UserService interface {
	RegisterUser(ctx context.Context, user data.RentalUser, password string) (data.RentalUser, error)
	GetUserById(ctx context.Context, id string) (data.RentalUser, error)
	GetUserByEmail(ctx context.Context, email string) (data.RentalUser, error)
	UpdateUser(ctx context.Context, user data.RentalUser) (data.RentalUser, error)
	DeleteUser(ctx context.Context, id string) error
	GetAllUsers(ctx context.Context) ([]data.RentalUser, error)
}

type userService struct {
	repo data.UserRepository
}

func (s userService) RegisterUser(ctx context.Context, user data.RentalUser, password string) (data.RentalUser, error) {
	if user.Email == "" || user.Name == "" || password == "" {
		return data.RentalUser{}, errors.New("ERROR: missing required user fields")
	}

	createdUser, err := s.repo.SaveUser(user)
	if err != nil {
		return data.RentalUser{}, fmt.Errorf("RegisterUser: %w", err)
	}
	return createdUser, nil
}

func (s userService) GetUserById(ctx context.Context, id string) (data.RentalUser, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return data.RentalUser{}, fmt.Errorf("GetUserById: invalid id format: %w", err)
	}
	user, err := s.repo.GetUserById(uid)
	if err != nil {
		return data.RentalUser{}, fmt.Errorf("GetUserById: %w", err)
	}
	return user, nil
}

func (s userService) GetUserByEmail(ctx context.Context, email string) (data.RentalUser, error) {
	if email == "" {
		return data.RentalUser{}, errors.New("ERROR: Email is empty")
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return data.RentalUser{}, fmt.Errorf("GetUserByEmail: %w", err)
	}
	return user, nil
}

func (s userService) UpdateUser(ctx context.Context, user data.RentalUser) (data.RentalUser, error) {
	if user.Id == uuid.Nil {
		return data.RentalUser{}, errors.New("ERROR: User id is empty")
	}

	updatedUser, err := s.repo.UpdateUserById(user.Id, user)
	if err != nil {
		return data.RentalUser{}, fmt.Errorf("UpdateUser: %w", err)
	}
	return updatedUser, nil
}

func (s userService) DeleteUser(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("DeleteUser: invalid id format: %w", err)
	}
	if err := s.repo.DeleteUserById(uid); err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}
	return nil
}

func (s userService) GetAllUsers(ctx context.Context) ([]data.RentalUser, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("GetAllUsers: %w", err)
	}

	return users, nil
}

func NewUserService(repo data.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
