package serviceTests

import (
	"context"
	"errors"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockCarRepository struct {
	saveCarFunc                     func(car data.Car) (data.Car, error)
	getCarByVinFunc                 func(vin string) (data.Car, error)
	updateCarFunc                   func(car data.Car) (data.Car, error)
	deleteCarByVinFunc              func(vin string) error
	getCarsAvailableInTimeRangeFunc func(startTime, endTime time.Time) ([]data.Car, error)
}

func (m *mockCarRepository) SaveCar(car data.Car) (data.Car, error) {
	if m.saveCarFunc != nil {
		return m.saveCarFunc(car)
	}
	return data.Car{}, nil
}

func (m *mockCarRepository) GetCarByVin(vin string) (data.Car, error) {
	if m.getCarByVinFunc != nil {
		return m.getCarByVinFunc(vin)
	}
	return data.Car{}, nil
}

func (m *mockCarRepository) UpdateCar(car data.Car) (data.Car, error) {
	if m.updateCarFunc != nil {
		return m.updateCarFunc(car)
	}
	return data.Car{}, nil
}

func (m *mockCarRepository) DeleteCarByVin(vin string) error {
	if m.deleteCarByVinFunc != nil {
		return m.deleteCarByVinFunc(vin)
	}
	return nil
}

func (m *mockCarRepository) GetCarsAvailableInTimeRange(startTime, endTime time.Time) ([]data.Car, error) {
	if m.getCarsAvailableInTimeRangeFunc != nil {
		return m.getCarsAvailableInTimeRangeFunc(startTime, endTime)
	}
	return nil, nil
}

// ====================
// Tests for CreateCar
// ====================
func TestCreateCar(t *testing.T) {
	tests := []struct {
		name          string
		inputCar      data.Car
		mockRepoFunc  func(car data.Car) (data.Car, error)
		expectedError string
	}{
		{
			name: "empty VIN",
			inputCar: data.Car{
				Vin:   "",
				Model: "Golf",
				Brand: "VW",
			},
			expectedError: "ERROR: Car vin is empty",
		},
		{
			name: "missing brand",
			inputCar: data.Car{
				Vin:   "1234ABC",
				Model: "Golf",
				Brand: "",
			},
			expectedError: "ERROR: Car model is empty",
		},
		{
			name: "repository error on save",
			inputCar: data.Car{
				Vin:   "1234ABC",
				Model: "Golf",
				Brand: "VW",
			},
			mockRepoFunc: func(car data.Car) (data.Car, error) {
				return data.Car{}, errors.New("database error")
			},
			expectedError: "database error",
		},
		{
			name: "success",
			inputCar: data.Car{
				Vin:         "1234ABC",
				Model:       "Golf",
				Brand:       "VW",
				ImageUrl:    "http://example.com/image.png",
				PricePerDay: "50",
			},
			mockRepoFunc: func(car data.Car) (data.Car, error) {
				// Simulate DB-Call
				return car, nil
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockCarRepository{
				saveCarFunc: tc.mockRepoFunc,
			}
			carSrv := service.NewCarService(mockRepo)

			result, err := carSrv.CreateCar(context.Background(), tc.inputCar)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.inputCar.Vin, result.Vin)
			assert.Equal(t, tc.inputCar.Brand, result.Brand)
			assert.Equal(t, tc.inputCar.Model, result.Model)
		})
	}
}

// ====================
// Tests for GetCarByVin
// ====================
func TestGetCarByVin(t *testing.T) {
	tests := []struct {
		name          string
		vin           string
		mockRepoFunc  func(vin string) (data.Car, error)
		expectedError string
		expectedCar   data.Car
	}{
		{
			name:          "empty VIN",
			vin:           "",
			expectedError: "ERROR: Car vin is empty",
		},
		{
			name: "repository error",
			vin:  "1234ABC",
			mockRepoFunc: func(vin string) (data.Car, error) {
				return data.Car{}, errors.New("repo error")
			},
			expectedError: "repo error",
		},
		{
			name: "success",
			vin:  "1234ABC",
			mockRepoFunc: func(vin string) (data.Car, error) {
				return data.Car{
					Vin:         "1234ABC",
					Model:       "Golf",
					Brand:       "VW",
					ImageUrl:    "http://example.com/img.png",
					PricePerDay: "50",
				}, nil
			},
			expectedError: "",
			expectedCar: data.Car{
				Vin:         "1234ABC",
				Model:       "Golf",
				Brand:       "VW",
				ImageUrl:    "http://example.com/img.png",
				PricePerDay: "50",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockCarRepository{
				getCarByVinFunc: tc.mockRepoFunc,
			}
			carSrv := service.NewCarService(mockRepo)
			result, err := carSrv.GetCarByVin(context.Background(), tc.vin)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCar.Vin, result.Vin)
			assert.Equal(t, tc.expectedCar.Model, result.Model)
			assert.Equal(t, tc.expectedCar.Brand, result.Brand)
		})
	}
}

// ====================
// Tests for UpdateCar
// ====================
func TestUpdateCar(t *testing.T) {
	tests := []struct {
		name          string
		inputCar      data.Car
		mockRepoFunc  func(car data.Car) (data.Car, error)
		expectedError string
	}{
		{
			name: "empty VIN",
			inputCar: data.Car{
				Vin:   "",
				Model: "Golf",
				Brand: "VW",
			},
			expectedError: "ERROR: Car vin is empty",
		},
		{
			name: "repository error on update",
			inputCar: data.Car{
				Vin:   "1234ABC",
				Model: "Golf",
				Brand: "VW",
			},
			mockRepoFunc: func(car data.Car) (data.Car, error) {
				return data.Car{}, errors.New("update error")
			},
			expectedError: "update error",
		},
		{
			name: "success",
			inputCar: data.Car{
				Vin:         "1234ABC",
				Model:       "Golf",
				Brand:       "VW",
				ImageUrl:    "http://example.com/image.png",
				PricePerDay: "50",
			},
			mockRepoFunc: func(car data.Car) (data.Car, error) {
				return car, nil
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockCarRepository{
				updateCarFunc: tc.mockRepoFunc,
			}
			carSrv := service.NewCarService(mockRepo)
			result, err := carSrv.UpdateCar(context.Background(), tc.inputCar)
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.inputCar.Vin, result.Vin)
			assert.Equal(t, tc.inputCar.Model, result.Model)
			assert.Equal(t, tc.inputCar.Brand, result.Brand)
		})
	}
}

// ====================
// Tests for DeleteCarByVin
// ====================
func TestDeleteCarByVin(t *testing.T) {
	tests := []struct {
		name          string
		vin           string
		mockRepoFunc  func(vin string) error
		expectedError string
	}{
		{
			name:          "empty VIN",
			vin:           "",
			expectedError: "ERROR: Car vin is empty",
		},
		{
			name: "repository error on delete",
			vin:  "1234ABC",
			mockRepoFunc: func(vin string) error {
				return errors.New("delete error")
			},
			expectedError: "delete error",
		},
		{
			name: "success",
			vin:  "1234ABC",
			mockRepoFunc: func(vin string) error {
				return nil
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockCarRepository{
				deleteCarByVinFunc: tc.mockRepoFunc,
			}
			carSrv := service.NewCarService(mockRepo)
			err := carSrv.DeleteCarByVin(context.Background(), tc.vin)
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
// Tests for GetCarsAvailableInTimeRange
// ====================
func TestGetCarsAvailableInTimeRange(t *testing.T) {
	start := time.Now()
	end := start.Add(24 * time.Hour)

	tests := []struct {
		name          string
		startTime     time.Time
		endTime       time.Time
		mockRepoFunc  func(startTime, endTime time.Time) ([]data.Car, error)
		expectedError string
		expectedCount int
	}{
		{
			name:          "endTime before startTime",
			startTime:     end,
			endTime:       start,
			expectedError: "EndTime is earlier than StartTime",
		},
		{
			name:      "repository error",
			startTime: start,
			endTime:   end,
			mockRepoFunc: func(startTime, endTime time.Time) ([]data.Car, error) {
				return nil, errors.New("range error")
			},
			expectedError: "range error",
		},
		{
			name:      "success with no cars",
			startTime: start,
			endTime:   end,
			mockRepoFunc: func(startTime, endTime time.Time) ([]data.Car, error) {
				return []data.Car{}, nil
			},
			expectedError: "",
			expectedCount: 0,
		},
		{
			name:      "success with cars",
			startTime: start,
			endTime:   end,
			mockRepoFunc: func(startTime, endTime time.Time) ([]data.Car, error) {
				return []data.Car{
					{
						Vin:         "1234ABC",
						Model:       "Golf",
						Brand:       "VW",
						ImageUrl:    "http://example.com/img.png",
						PricePerDay: "50",
					},
					{
						Vin:         "5678DEF",
						Model:       "Polo",
						Brand:       "VW",
						ImageUrl:    "http://example.com/img2.png",
						PricePerDay: "40",
					},
				}, nil
			},
			expectedError: "",
			expectedCount: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockCarRepository{
				getCarsAvailableInTimeRangeFunc: tc.mockRepoFunc,
			}
			carSrv := service.NewCarService(mockRepo)
			result, err := carSrv.GetCarsAvailableInTimeRange(context.Background(), tc.startTime, tc.endTime)
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
