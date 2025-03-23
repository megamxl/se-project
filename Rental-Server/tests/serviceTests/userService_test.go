package serviceTests

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/megamxl/se-project/Rental-Server/api/DTO"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockUserRepository struct {
	saveUserFunc          func(user data.RentalUser) (data.RentalUser, error)
	getUserByEmailFunc    func(email string) (data.RentalUser, error)
	getUserByIdFunc       func(id uuid.UUID) (data.RentalUser, error)
	updateUserByIdFunc    func(id uuid.UUID, update data.RentalUser) (data.RentalUser, error)
	deleteUserByIdFunc    func(id uuid.UUID) error
	getAllUsersFunc       func() ([]data.RentalUser, error)
	updateUserByEmailFunc func(email string, update data.RentalUser) (data.RentalUser, error)
}

func (m *mockUserRepository) UpdateUserByEmail(email string, update data.RentalUser) (data.RentalUser, error) {
	if m.updateUserByEmailFunc != nil {
		return m.updateUserByEmailFunc(email, update)
	}
	return data.RentalUser{}, nil
}

func (m *mockUserRepository) SaveUser(user data.RentalUser) (data.RentalUser, error) {
	if m.saveUserFunc != nil {
		return m.saveUserFunc(user)
	}
	return data.RentalUser{}, nil
}

func (m *mockUserRepository) GetUserByEmail(email string) (data.RentalUser, error) {
	if m.getUserByEmailFunc != nil {
		return m.getUserByEmailFunc(email)
	}
	return data.RentalUser{}, nil
}

func (m *mockUserRepository) GetUserById(id uuid.UUID) (data.RentalUser, error) {
	if m.getUserByIdFunc != nil {
		return m.getUserByIdFunc(id)
	}
	return data.RentalUser{}, nil
}

func (m *mockUserRepository) UpdateUserById(id uuid.UUID, update data.RentalUser) (data.RentalUser, error) {
	if m.updateUserByIdFunc != nil {
		return m.updateUserByIdFunc(id, update)
	}
	return data.RentalUser{}, nil
}

func (m *mockUserRepository) DeleteUserById(id uuid.UUID) error {
	if m.deleteUserByIdFunc != nil {
		return m.deleteUserByIdFunc(id)
	}
	return nil
}

func (m *mockUserRepository) GetAllUsers() ([]data.RentalUser, error) {
	if m.getAllUsersFunc != nil {
		return m.getAllUsersFunc()
	}
	return nil, nil
}

// ====================
// Test for RegisterUser
// ====================
func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name          string
		inputUser     DTO.RentalUser
		password      string
		mockFunc      func(user data.RentalUser) (data.RentalUser, error)
		expectedError string
	}{
		{
			name: "missing fields",
			inputUser: DTO.RentalUser{
				Email: "",
				Name:  "Alice",
			},
			password:      "secret",
			expectedError: "ERROR: missing required user fields",
		},
		{
			name: "repository error",
			inputUser: DTO.RentalUser{
				Email: "alice@example.com",
				Name:  "Alice",
			},
			password: "secret",
			mockFunc: func(user data.RentalUser) (data.RentalUser, error) {
				return data.RentalUser{}, errors.New("save error")
			},
			expectedError: "save error",
		},
		{
			name: "success",
			inputUser: DTO.RentalUser{
				Email: "alice@example.com",
				Name:  "Alice",
			},
			password: "secret",
			mockFunc: func(user data.RentalUser) (data.RentalUser, error) {
				user.Id = uuid.New()
				return user, nil
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockUserRepository{
				saveUserFunc: tc.mockFunc,
			}
			userSrv := service.NewUserService(mockRepo)
			result, err := userSrv.RegisterUser(context.Background(), tc.inputUser, tc.password)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.NotEqual(t, uuid.Nil, result.Id)
			assert.Equal(t, tc.inputUser.Email, result.Email)
			assert.Equal(t, tc.inputUser.Name, result.Name)
		})
	}
}

// ====================
// Test for GetUserById
// ====================
func TestGetUserById(t *testing.T) {
	validID := uuid.New()
	tests := []struct {
		name          string
		inputID       string
		mockFunc      func(id uuid.UUID) (data.RentalUser, error)
		expectedError string
		expectedUser  DTO.RentalUser
	}{
		{
			name:          "invalid id format",
			inputID:       "not-a-uuid",
			expectedError: "invalid id format",
		},
		{
			name:    "repository error",
			inputID: validID.String(),
			mockFunc: func(id uuid.UUID) (data.RentalUser, error) {
				return data.RentalUser{}, errors.New("get error")
			},
			expectedError: "get error",
		},
		{
			name:    "success",
			inputID: validID.String(),
			mockFunc: func(id uuid.UUID) (data.RentalUser, error) {
				return data.RentalUser{
					Id:       id,
					Email:    "bob@example.com",
					Name:     "Bob",
					Password: "hashed:secret",
				}, nil
			},
			expectedError: "",
			expectedUser: DTO.RentalUser{
				Id:    validID,
				Email: "bob@example.com",
				Name:  "Bob",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockUserRepository{
				getUserByIdFunc: tc.mockFunc,
			}
			userSrv := service.NewUserService(mockRepo)
			result, err := userSrv.GetUserById(context.Background(), tc.inputID)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedUser.Id, result.Id)
			assert.Equal(t, tc.expectedUser.Email, result.Email)
			assert.Equal(t, tc.expectedUser.Name, result.Name)
		})
	}
}

// ====================
// Test for GetUserByEmail
// ====================
func TestGetUserByEmail(t *testing.T) {
	tests := []struct {
		name          string
		inputEmail    string
		mockFunc      func(email string) (data.RentalUser, error)
		expectedError string
		expectedUser  DTO.RentalUser
	}{
		{
			name:          "empty email",
			inputEmail:    "",
			expectedError: "ERROR: Email is empty",
		},
		{
			name:       "repository error",
			inputEmail: "charlie@example.com",
			mockFunc: func(email string) (data.RentalUser, error) {
				return data.RentalUser{}, errors.New("email error")
			},
			expectedError: "email error",
		},
		{
			name:       "success",
			inputEmail: "charlie@example.com",
			mockFunc: func(email string) (data.RentalUser, error) {
				return data.RentalUser{
					Id:       uuid.New(),
					Email:    email,
					Name:     "Charlie",
					Password: "hashed:secret",
				}, nil
			},
			expectedError: "",
			expectedUser: DTO.RentalUser{
				Email: "charlie@example.com",
				Name:  "Charlie",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockUserRepository{
				getUserByEmailFunc: tc.mockFunc,
			}
			userSrv := service.NewUserService(mockRepo)
			result, err := userSrv.GetUserByEmail(context.Background(), tc.inputEmail)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedUser.Email, result.Email)
			assert.Equal(t, tc.expectedUser.Name, result.Name)
		})
	}
}

// ====================
// Test for UpdateUser
// ====================
func TestUpdateUser(t *testing.T) {
	validID := uuid.New()
	tests := []struct {
		name          string
		inputUser     DTO.RentalUser
		mockFunc      func(id uuid.UUID, update data.RentalUser) (data.RentalUser, error)
		expectedError string
	}{
		{
			name: "empty user id",
			inputUser: DTO.RentalUser{
				Id:       uuid.Nil,
				Email:    "dave@example.com",
				Name:     "Dave",
				Password: "hashed:secret",
			},
			expectedError: "ERROR: User id is empty",
		},
		{
			name: "repository error",
			inputUser: DTO.RentalUser{
				Id:       validID,
				Email:    "dave@example.com",
				Name:     "Dave",
				Password: "hashed:secret",
			},
			mockFunc: func(id uuid.UUID, update data.RentalUser) (data.RentalUser, error) {
				return data.RentalUser{}, errors.New("update failed")
			},
			expectedError: "update failed",
		},
		{
			name: "success",
			inputUser: DTO.RentalUser{
				Id:       validID,
				Email:    "dave@example.com",
				Name:     "Dave",
				Password: "hashed:secret",
			},
			mockFunc: func(id uuid.UUID, update data.RentalUser) (data.RentalUser, error) {
				return update, nil
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockUserRepository{
				updateUserByIdFunc: tc.mockFunc,
			}
			userSrv := service.NewUserService(mockRepo)
			result, err := userSrv.UpdateUser(context.Background(), tc.inputUser)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.inputUser.Id, result.Id)
			assert.Equal(t, tc.inputUser.Email, result.Email)
			assert.Equal(t, tc.inputUser.Name, result.Name)
		})
	}
}

// ====================
// Test for DeleteUser
// ====================
func TestDeleteUser(t *testing.T) {
	validID := uuid.New()
	tests := []struct {
		name          string
		inputID       string
		mockFunc      func(id uuid.UUID) error
		expectedError string
	}{
		{
			name:          "invalid id format",
			inputID:       "not-a-uuid",
			expectedError: "invalid id format",
		},
		{
			name:    "repository error",
			inputID: validID.String(),
			mockFunc: func(id uuid.UUID) error {
				return errors.New("delete failed")
			},
			expectedError: "delete failed",
		},
		{
			name:    "success",
			inputID: validID.String(),
			mockFunc: func(id uuid.UUID) error {
				return nil
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockUserRepository{
				deleteUserByIdFunc: tc.mockFunc,
			}
			userSrv := service.NewUserService(mockRepo)
			err := userSrv.DeleteUser(context.Background(), tc.inputID)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
		})
	}
}

// ====================
// Test for GetAllUsers
// ====================
func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name          string
		mockFunc      func() ([]data.RentalUser, error)
		expectedError string
		expectedCount int
	}{
		{
			name: "repository error",
			mockFunc: func() ([]data.RentalUser, error) {
				return nil, errors.New("get all error")
			},
			expectedError: "get all error",
		},
		{
			name: "empty list",
			mockFunc: func() ([]data.RentalUser, error) {
				return []data.RentalUser{}, nil
			},
			expectedError: "",
			expectedCount: 0,
		},
		{
			name: "success",
			mockFunc: func() ([]data.RentalUser, error) {
				return []data.RentalUser{
					{
						Id:       uuid.New(),
						Email:    "user1@example.com",
						Name:     "User One",
						Password: "hashed:pass1",
					},
					{
						Id:       uuid.New(),
						Email:    "user2@example.com",
						Name:     "User Two",
						Password: "hashed:pass2",
					},
				}, nil
			},
			expectedError: "",
			expectedCount: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockUserRepository{
				getAllUsersFunc: tc.mockFunc,
			}
			userSrv := service.NewUserService(mockRepo)
			result, err := userSrv.GetAllUsers(context.Background())
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCount, len(result))
		})
	}
}
