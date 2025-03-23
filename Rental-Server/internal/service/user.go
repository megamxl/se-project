package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/megamxl/se-project/Rental-Server/api/DTO"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
)

type UserService interface {
	RegisterUser(ctx context.Context, user DTO.RentalUser, password string) (DTO.RentalUser, error)
	GetUserById(ctx context.Context, id string) (DTO.RentalUser, error)
	GetUserByEmail(ctx context.Context, email string) (DTO.RentalUser, error)
	UpdateUser(ctx context.Context, user DTO.RentalUser) (DTO.RentalUser, error)
	DeleteUser(ctx context.Context, id string) error
	GetAllUsers(ctx context.Context) ([]DTO.RentalUser, error)
}

type userService struct {
	repo data.UserRepository
}

func (s userService) RegisterUser(ctx context.Context, user DTO.RentalUser, password string) (DTO.RentalUser, error) {
	if user.Email == "" || user.Name == "" || password == "" {
		return DTO.RentalUser{}, errors.New("ERROR: missing required user fields")
	}

	dataUser := MapDTOUserToDataUser(user)

	createdUser, err := s.repo.SaveUser(dataUser)
	if err != nil {
		return DTO.RentalUser{}, fmt.Errorf("RegisterUser: %w", err)
	}
	return MapDataUserToDTOUser(createdUser), nil
}

func (s userService) GetUserById(ctx context.Context, id string) (DTO.RentalUser, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return DTO.RentalUser{}, fmt.Errorf("GetUserById: invalid id format: %w", err)
	}
	dataUser, err := s.repo.GetUserById(uid)
	if err != nil {
		return DTO.RentalUser{}, fmt.Errorf("GetUserById: %w", err)
	}
	return MapDataUserToDTOUser(dataUser), nil
}

func (s userService) GetUserByEmail(ctx context.Context, email string) (DTO.RentalUser, error) {
	if email == "" {
		return DTO.RentalUser{}, errors.New("ERROR: Email is empty")
	}
	dataUser, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return DTO.RentalUser{}, fmt.Errorf("GetUserByEmail: %w", err)
	}
	return MapDataUserToDTOUser(dataUser), nil
}

func (s userService) UpdateUser(ctx context.Context, user DTO.RentalUser) (DTO.RentalUser, error) {
	if user.Id == uuid.Nil {
		return DTO.RentalUser{}, errors.New("ERROR: User id is empty")
	}
	dataUser := MapDTOUserToDataUser(user)

	updatedUser, err := s.repo.UpdateUserById(dataUser.Id, dataUser)
	if err != nil {
		return DTO.RentalUser{}, fmt.Errorf("UpdateUser: %w", err)
	}
	return MapDataUserToDTOUser(updatedUser), nil
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

func (s userService) GetAllUsers(ctx context.Context) ([]DTO.RentalUser, error) {
	dataUsers, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("GetAllUsers: %w", err)
	}
	dtoUsers := make([]DTO.RentalUser, len(dataUsers))
	for index, du := range dataUsers {
		dtoUsers[index] = MapDataUserToDTOUser(du)
	}
	return dtoUsers, nil
}

func NewUserService(repo data.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func MapDataUserToDTOUser(dataUser data.RentalUser) DTO.RentalUser {
	return DTO.RentalUser{
		Id:       dataUser.Id,
		Name:     dataUser.Name,
		Email:    dataUser.Email,
		Password: dataUser.Password,
	}
}

func MapDTOUserToDataUser(dtoUser DTO.RentalUser) data.RentalUser {
	return data.RentalUser{
		Id:       dtoUser.Id,
		Name:     dtoUser.Name,
		Email:    dtoUser.Email,
		Password: dtoUser.Password,
	}
}
