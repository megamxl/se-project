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

type mockBookingRepository struct {
	saveBookingFunc          func(booking data.Booking) (data.Booking, error)
	getBookingByIdFunc       func(id uuid.UUID) (data.Booking, error)
	deleteBookingByIdFunc    func(id uuid.UUID) error
	getAllBookingsByUserFunc func(userId uuid.UUID) ([]data.Booking, error)
	getAllBookingsFunc       func() ([]data.Booking, error)
	getBookingsByVinFunc     func(vin string) (data.Booking, error)
	deleteBookingsByVinFunc  func(vin string) error
}

func (m *mockBookingRepository) GetBookingsByVin(vin string) (data.Booking, error) {
	if m.getBookingsByVinFunc != nil {
		return m.getBookingsByVinFunc(vin)
	}
	return data.Booking{}, nil
}

func (m *mockBookingRepository) DeleteBookingsByVin(vin string) error {
	if m.deleteBookingsByVinFunc != nil {
		return m.deleteBookingsByVinFunc(vin)
	}
	return nil
}

func (m *mockBookingRepository) SaveBooking(booking data.Booking) (data.Booking, error) {
	if m.saveBookingFunc != nil {
		return m.saveBookingFunc(booking)
	}
	return data.Booking{}, nil
}

func (m *mockBookingRepository) GetBookingById(id uuid.UUID) (data.Booking, error) {
	if m.getBookingByIdFunc != nil {
		return m.getBookingByIdFunc(id)
	}
	return data.Booking{}, nil
}

func (m *mockBookingRepository) DeleteBookingById(id uuid.UUID) error {
	if m.deleteBookingByIdFunc != nil {
		return m.deleteBookingByIdFunc(id)
	}
	return nil
}

func (m *mockBookingRepository) GetAllBookingsByUser(userId uuid.UUID) ([]data.Booking, error) {
	if m.getAllBookingsByUserFunc != nil {
		return m.getAllBookingsByUserFunc(userId)
	}
	return nil, nil
}

func (m *mockBookingRepository) GetAllBookings() ([]data.Booking, error) {
	if m.getAllBookingsFunc != nil {
		return m.getAllBookingsFunc()
	}
	return nil, nil
}

// =================
// Tests for BookCar
// =================
func TestBookCar(t *testing.T) {
	tests := []struct {
		name          string
		inputBooking  DTO.Booking
		mockSaveFunc  func(booking data.Booking) (data.Booking, error)
		expectedError string
	}{
		{
			name: "empty VIN",
			inputBooking: DTO.Booking{
				CarVin: "",
				Status: "",
			},
			expectedError: "BookCar: VIN is empty",
		},
		{
			name: "repository error on save",
			inputBooking: DTO.Booking{
				CarVin: "ABC123",
				Status: "",
			},
			mockSaveFunc: func(booking data.Booking) (data.Booking, error) {
				return data.Booking{}, errors.New("repo save error")
			},
			expectedError: "repo save error",
		},
		{
			name: "success",
			inputBooking: DTO.Booking{
				CarVin: "ABC123",
				Status: "",
			},
			mockSaveFunc: func(booking data.Booking) (data.Booking, error) {
				booking.Id = uuid.New()
				return booking, nil
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockBookingRepository{
				saveBookingFunc: tc.mockSaveFunc,
			}
			bookSrv := service.NewBookingService(mockRepo)
			result, err := bookSrv.BookCar(context.Background(), tc.inputBooking)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, "pending", result.Status)
			assert.Equal(t, tc.inputBooking.CarVin, result.CarVin)
		})
	}
}

// =======================
// Tests for GetBookingById
// =======================
func TestGetBookingById(t *testing.T) {
	validID := uuid.New().String()
	tests := []struct {
		name            string
		inputID         string
		mockGetFunc     func(id uuid.UUID) (data.Booking, error)
		expectedError   string
		expectedBooking DTO.Booking
	}{
		{
			name:          "invalid id format",
			inputID:       "not-a-uuid",
			expectedError: "invalid id format",
		},
		{
			name:    "repository error",
			inputID: validID,
			mockGetFunc: func(id uuid.UUID) (data.Booking, error) {
				return data.Booking{}, errors.New("get error")
			},
			expectedError: "get error",
		},
		{
			name:    "success",
			inputID: validID,
			mockGetFunc: func(id uuid.UUID) (data.Booking, error) {
				return data.Booking{
					Id:     id,
					CarVin: "ABC123",
					Status: "pending",
				}, nil
			},
			expectedError: "",
			expectedBooking: DTO.Booking{
				Id:     uuid.MustParse(validID),
				CarVin: "ABC123",
				Status: "pending",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockBookingRepository{
				getBookingByIdFunc: tc.mockGetFunc,
			}
			bookSrv := service.NewBookingService(mockRepo)
			result, err := bookSrv.GetBookingById(context.Background(), tc.inputID)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBooking.CarVin, result.CarVin)
			assert.Equal(t, tc.expectedBooking.Status, result.Status)
			assert.Equal(t, tc.inputID, result.Id.String())
		})
	}
}

// =======================
// Tests for UpdateBooking
// =======================
func TestUpdateBooking(t *testing.T) {
	validID := uuid.New().String()
	tests := []struct {
		name          string
		bookingID     string
		newStatus     string
		mockGetFunc   func(id uuid.UUID) (data.Booking, error)
		mockSaveFunc  func(booking data.Booking) (data.Booking, error)
		expectedError string
	}{
		{
			name:          "invalid booking id format",
			bookingID:     "invalid-uuid",
			newStatus:     "confirmed",
			expectedError: "invalid booking id",
		},
		{
			name:      "repository error on get",
			bookingID: validID,
			newStatus: "confirmed",
			mockGetFunc: func(id uuid.UUID) (data.Booking, error) {
				return data.Booking{}, errors.New("get error")
			},
			expectedError: "get error",
		},
		{
			name:      "repository error on save",
			bookingID: validID,
			newStatus: "confirmed",
			mockGetFunc: func(id uuid.UUID) (data.Booking, error) {
				return data.Booking{
					Id:     id,
					CarVin: "ABC123",
					Status: "pending",
				}, nil
			},
			mockSaveFunc: func(booking data.Booking) (data.Booking, error) {
				return data.Booking{}, errors.New("save error")
			},
			expectedError: "save error",
		},
		{
			name:      "success",
			bookingID: validID,
			newStatus: "confirmed",
			mockGetFunc: func(id uuid.UUID) (data.Booking, error) {
				return data.Booking{
					Id:     id,
					CarVin: "ABC123",
					Status: "pending",
				}, nil
			},
			mockSaveFunc: func(booking data.Booking) (data.Booking, error) {
				booking.Status = "confirmed"
				return booking, nil
			},
			expectedError: "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockBookingRepository{
				getBookingByIdFunc: tc.mockGetFunc,
				saveBookingFunc:    tc.mockSaveFunc,
			}
			bookSrv := service.NewBookingService(mockRepo)
			result, err := bookSrv.UpdateBooking(context.Background(), tc.bookingID, tc.newStatus)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.newStatus, result.Status)
		})
	}
}

// =======================
// Tests for DeleteBooking
// =======================
func TestDeleteBooking(t *testing.T) {
	validID := uuid.New().String()
	tests := []struct {
		name           string
		bookingID      string
		mockDeleteFunc func(id uuid.UUID) error
		expectedError  string
	}{
		{
			name:          "invalid booking id",
			bookingID:     "not-a-uuid",
			expectedError: "invalid booking id",
		},
		{
			name:      "repository error",
			bookingID: validID,
			mockDeleteFunc: func(id uuid.UUID) error {
				return errors.New("delete error")
			},
			expectedError: "delete error",
		},
		{
			name:      "success",
			bookingID: validID,
			mockDeleteFunc: func(id uuid.UUID) error {
				return nil
			},
			expectedError: "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockBookingRepository{
				deleteBookingByIdFunc: tc.mockDeleteFunc,
			}
			bookSrv := service.NewBookingService(mockRepo)
			err := bookSrv.DeleteBooking(context.Background(), tc.bookingID)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
		})
	}
}

// ==============================
// Tests for GetAllBookingsByUser
// ==============================
func TestGetAllBookingsByUser(t *testing.T) {
	validUserID := uuid.New().String()
	tests := []struct {
		name          string
		userID        string
		mockFunc      func(userId uuid.UUID) ([]data.Booking, error)
		expectedError string
		expectedCount int
	}{
		{
			name:          "invalid user id",
			userID:        "invalid-uuid",
			expectedError: "invalid user id",
		},
		{
			name:   "repository error",
			userID: validUserID,
			mockFunc: func(userId uuid.UUID) ([]data.Booking, error) {
				return nil, errors.New("user bookings error")
			},
			expectedError: "user bookings error",
		},
		{
			name:   "success with empty list",
			userID: validUserID,
			mockFunc: func(userId uuid.UUID) ([]data.Booking, error) {
				return []data.Booking{}, nil
			},
			expectedError: "",
			expectedCount: 0,
		},
		{
			name:   "success with bookings",
			userID: validUserID,
			mockFunc: func(userId uuid.UUID) ([]data.Booking, error) {
				return []data.Booking{
					{
						Id:     uuid.New(),
						CarVin: "ABC123",
						Status: "pending",
					},
					{
						Id:     uuid.New(),
						CarVin: "DEF456",
						Status: "confirmed",
					},
				}, nil
			},
			expectedError: "",
			expectedCount: 2,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockBookingRepository{
				getAllBookingsByUserFunc: tc.mockFunc,
			}
			bookSrv := service.NewBookingService(mockRepo)
			result, err := bookSrv.GetAllBookingsByUser(context.Background(), tc.userID)
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

// ========================
// Tests for GetAllBookings
// ========================
func TestGetAllBookings(t *testing.T) {
	tests := []struct {
		name          string
		mockFunc      func() ([]data.Booking, error)
		expectedError string
		expectedCount int
	}{
		{
			name: "repository error",
			mockFunc: func() ([]data.Booking, error) {
				return nil, errors.New("all bookings error")
			},
			expectedError: "all bookings error",
		},
		{
			name: "success with empty list",
			mockFunc: func() ([]data.Booking, error) {
				return []data.Booking{}, nil
			},
			expectedError: "",
			expectedCount: 0,
		},
		{
			name: "success with bookings",
			mockFunc: func() ([]data.Booking, error) {
				return []data.Booking{
					{
						Id:     uuid.New(),
						CarVin: "ABC123",
						Status: "pending",
					},
					{
						Id:     uuid.New(),
						CarVin: "DEF456",
						Status: "confirmed",
					},
				}, nil
			},
			expectedError: "",
			expectedCount: 2,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockBookingRepository{
				getAllBookingsFunc: tc.mockFunc,
			}
			bookSrv := service.NewBookingService(mockRepo)
			result, err := bookSrv.GetAllBookings(context.Background())
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
