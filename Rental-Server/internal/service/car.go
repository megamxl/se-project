package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/megamxl/se-project/Rental-Server/api/DTO"
	"time"

	"github.com/megamxl/se-project/Rental-Server/internal/data"
)

type CarService interface {
	GetCarByVin(ctx context.Context, vin string) (DTO.Car, error)
	CreateCar(ctx context.Context, car DTO.Car) (DTO.Car, error)
	UpdateCar(ctx context.Context, car DTO.Car) (DTO.Car, error)
	DeleteCarByVin(ctx context.Context, vin string) error
	GetCarsAvailableInTimeRange(ctx context.Context, startTime, endTime time.Time) ([]DTO.Car, error)
}

type carService struct {
	repo data.CarRepository
}

func (s *carService) GetCarByVin(ctx context.Context, vin string) (DTO.Car, error) {
	if vin == "" {
		return DTO.Car{}, errors.New("ERROR: Car vin is empty")
	}

	dbCar, err := s.repo.GetCarByVin(vin)
	if err != nil {
		return DTO.Car{}, fmt.Errorf("GetCarByVin: %w", err)
	}

	return MapDataCarToDTOCar(dbCar), nil
}

func (s *carService) CreateCar(ctx context.Context, car DTO.Car) (DTO.Car, error) {
	if car.Vin == "" {
		return DTO.Car{}, errors.New("ERROR: Car vin is empty")
	}
	if car.Brand == "" || car.Model == "" {
		return DTO.Car{}, errors.New("ERROR: Car model is empty")
	}

	dataCar := MapDTOCarToDataCar(car)

	createdCar, err := s.repo.SaveCar(dataCar)
	if err != nil {
		return DTO.Car{}, fmt.Errorf("CreateCar: %w", err)
	}

	return MapDataCarToDTOCar(createdCar), nil
}

func (s *carService) UpdateCar(ctx context.Context, car DTO.Car) (DTO.Car, error) {
	if car.Vin == "" {
		return DTO.Car{}, errors.New("ERROR: Car vin is empty")
	}

	dataCar := MapDTOCarToDataCar(car)

	updatedCar, err := s.repo.UpdateCar(dataCar)
	if err != nil {
		return DTO.Car{}, fmt.Errorf("UpdateCar: %w", err)
	}

	return MapDataCarToDTOCar(updatedCar), nil
}

func (s *carService) DeleteCarByVin(ctx context.Context, vin string) error {
	if vin == "" {
		return errors.New("ERROR: Car vin is empty")
	}

	err := s.repo.DeleteCarByVin(vin)
	if err != nil {
		return fmt.Errorf("DeleteCarByVin: %w", err)
	}

	return nil
}

func (s *carService) GetCarsAvailableInTimeRange(ctx context.Context, startTime, endTime time.Time) ([]DTO.Car, error) {
	if endTime.Before(startTime) {
		return nil, errors.New("EndTime is earlier than StartTime")
	}

	dataCars, err := s.repo.GetCarsAvailableInTimeRange(startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("GetCarsAvailableInTimeRange: %w", err)
	}

	dtoCars := make([]DTO.Car, len(dataCars))
	for index, dataCar := range dataCars {
		dtoCars[index] = MapDataCarToDTOCar(dataCar)
	}

	return dtoCars, nil
}

func NewCarService(repo data.CarRepository) CarService {
	return &carService{
		repo: repo,
	}
}

func MapDataCarToDTOCar(dataCar data.Car) DTO.Car {
	return DTO.Car{
		Vin:         dataCar.Vin,
		Model:       dataCar.Model,
		Brand:       dataCar.Brand,
		ImageUrl:    dataCar.ImageUrl,
		PricePerDay: dataCar.PricePerDay,
	}
}

func MapDTOCarToDataCar(dtoCar DTO.Car) data.Car {
	return data.Car{
		Vin:         dtoCar.Vin,
		Model:       dtoCar.Model,
		Brand:       dtoCar.Brand,
		ImageUrl:    dtoCar.ImageUrl,
		PricePerDay: dtoCar.PricePerDay,
	}
}
