package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/megamxl/se-project/Rental-Server/internal/data"
)

type CarService interface {
	GetCarByVin(ctx context.Context, vin string) (data.Car, error)
	CreateCar(ctx context.Context, car data.Car) (data.Car, error)
	UpdateCar(ctx context.Context, car data.Car) (data.Car, error)
	DeleteCarByVin(ctx context.Context, vin string) error
	GetCarsAvailableInTimeRange(ctx context.Context, startTime, endTime time.Time) ([]data.Car, error)
}

type carService struct {
	repo data.CarRepository
}

func (s *carService) GetCarByVin(ctx context.Context, vin string) (data.Car, error) {
	if vin == "" {
		return data.Car{}, errors.New("ERROR: Car vin is empty")
	}

	dbCar, err := s.repo.GetCarByVin(vin)
	if err != nil {
		return data.Car{}, fmt.Errorf("GetCarByVin: %w", err)
	}

	return dbCar, nil
}

func (s *carService) CreateCar(ctx context.Context, car data.Car) (data.Car, error) {
	if car.Vin == "" {
		return data.Car{}, errors.New("ERROR: Car vin is empty")
	}
	if car.Brand == "" || car.Model == "" {
		return data.Car{}, errors.New("ERROR: Car model is empty")
	}

	createdCar, err := s.repo.SaveCar(car)
	if err != nil {
		return data.Car{}, fmt.Errorf("CreateCar: %w", err)
	}

	return createdCar, nil
}

func (s *carService) UpdateCar(ctx context.Context, car data.Car) (data.Car, error) {
	if car.Vin == "" {
		return data.Car{}, errors.New("ERROR: Car vin is empty")
	}

	updatedCar, err := s.repo.UpdateCar(car)
	if err != nil {
		return data.Car{}, fmt.Errorf("UpdateCar: %w", err)
	}

	return updatedCar, nil
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

func (s *carService) GetCarsAvailableInTimeRange(ctx context.Context, startTime, endTime time.Time) ([]data.Car, error) {
	if endTime.Before(startTime) {
		return nil, errors.New("EndTime is earlier than StartTime")
	}

	dataCars, err := s.repo.GetCarsAvailableInTimeRange(startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("GetCarsAvailableInTimeRange: %w", err)
	}

	return dataCars, nil
}

func NewCarService(repo data.CarRepository) CarService {
	return &carService{
		repo: repo,
	}
}
