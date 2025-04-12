package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/carEvents"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/converter"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/megamxl/se-project/Rental-Server/internal/data"
)

type CarService interface {
	GetCarByVin(ctx context.Context, vin string) (data.Car, error)
	CreateCar(ctx context.Context, car data.Car) (data.Car, error)
	UpdateCar(ctx context.Context, car data.Car) (data.Car, error)
	DeleteCarByVin(ctx context.Context, vin string) error
	GetCarsAvailableInTimeRange(ctx context.Context, startTime, endTime time.Time, currency string) ([]data.Car, error)
}

type carService struct {
	repo     data.CarRepository
	conv     converter.Converter
	producer carEvents.CarEvent
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

	if s.producer != nil {
		s.producer.AddCar(createdCar)
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

	if s.producer != nil {
		s.producer.UpdateCar(updatedCar)
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

	if s.producer != nil {
		s.producer.RemoveCar(data.Car{Vin: vin})
	}

	return nil
}

func (s *carService) GetCarsAvailableInTimeRange(ctx context.Context, startTime, endTime time.Time, currency string) ([]data.Car, error) {
	if endTime.Before(startTime) {
		return nil, errors.New("EndTime is earlier than StartTime")
	}

	dataCars, err := s.repo.GetCarsAvailableInTimeRange(startTime, endTime)

	for i, dataCar := range dataCars {
		convert, err := s.conv.Convert(converter.Request{
			GivenCurrency:  "USD",
			Amount:         dataCar.PricePerDay,
			TargetCurrency: currency,
		})
		if err != nil {
			slog.Error("Conversion Error:", err)
			continue
		}

		dataCars[i].PricePerDay = convert.Amount
	}

	if err != nil {
		return nil, fmt.Errorf("GetCarsAvailableInTimeRange: %w", err)
	}

	return dataCars, nil
}

func NewCarService(repo data.CarRepository, conv converter.Converter) CarService {

	service := carService{
		repo: repo,
		conv: conv,
	}

	if os.Getenv("PULSAR_PRODUCER") == "true" {
		client, err := pulsar.NewClient(pulsar.ClientOptions{
			URL:               os.Getenv("PULSAR_URL"),
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
		if err != nil {
			slog.Error("Could not instantiate Pulsar client: %v", err)
		}

		producer, err := client.CreateProducer(pulsar.ProducerOptions{
			Topic: "car-events",
		})

		if err != nil {
			log.Fatalf("Could not instantiate Pulsar producer: %v", err)
		}

		service.producer = carEvents.PulsarProducer{Producer: producer}
	}

	return &service
}
