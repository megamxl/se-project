package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/converter"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"log/slog"
	"time"
)

type BookingService interface {
	BookCar(ctx context.Context, req data.Booking, currency string, pricePerDayInUSD float64) (data.Booking, error)
	GetBookingById(ctx context.Context, id string) (data.Booking, error)
	UpdateBooking(ctx context.Context, bookingId string, status string) (data.Booking, error)
	DeleteBooking(ctx context.Context, bookingId string) error
	GetAllBookingsByUser(ctx context.Context, userId string) ([]data.Booking, error)
	GetAllBookings(ctx context.Context) ([]data.Booking, error)
	GetAllBookingsInTimeRange(from, to time.Time) ([]data.Booking, error)
}

type bookingService struct {
	repo data.BookingRepository
	conv converter.Converter
}

func (s bookingService) GetAllBookingsInTimeRange(from, to time.Time) ([]data.Booking, error) {

	timeRange, err := s.repo.GetBookingsInTimeRange(from, to)
	if err != nil {
		return nil, err
	}

	return timeRange, nil
}

func (s bookingService) BookCar(ctx context.Context, req data.Booking, currency string, pricePerDayInUSD float64) (data.Booking, error) {
	if req.CarVin == "" {
		return data.Booking{}, errors.New("BookCar: VIN is empty")
	}

	if req.Status == "" {
		req.Status = "pending"
	}

	between := GetDurationBetween(req.StartTime, req.EndTime).Hours() / 24

	convert, err := s.conv.Convert(converter.Request{
		GivenCurrency:  "USD",
		Amount:         pricePerDayInUSD,
		TargetCurrency: currency,
	})

	req.Currency = currency
	req.AmountPaid = convert.Amount * between

	if err != nil {
		slog.Error(" conversion error creating Booking in USD", err)
		req.Currency = "USD"
		req.AmountPaid = pricePerDayInUSD * between
	}

	saved, err := s.repo.SaveBooking(req)
	if err != nil {
		return data.Booking{}, fmt.Errorf("BookCar: %w", err)
	}

	return saved, nil
}

func (s bookingService) GetBookingById(ctx context.Context, id string) (data.Booking, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return data.Booking{}, fmt.Errorf("GetBookingById: invalid id format: %w", err)
	}
	booking, err := s.repo.GetBookingById(uid)
	if err != nil {
		return data.Booking{}, fmt.Errorf("GetBookingById: %w", err)
	}
	return booking, nil
}

func (s bookingService) UpdateBooking(ctx context.Context, bookingId string, status string) (data.Booking, error) {
	uid, err := uuid.Parse(bookingId)
	if err != nil {
		return data.Booking{}, fmt.Errorf("UpdateBooking: invalid booking id: %w", err)
	}

	booking, err := s.repo.GetBookingById(uid)
	if err != nil {
		return data.Booking{}, fmt.Errorf("UpdateBooking: %w", err)
	}
	booking.Status = status

	updatedBooking, err := s.repo.SaveBooking(booking)
	if err != nil {
		return data.Booking{}, fmt.Errorf("UpdateBooking: %w", err)
	}
	return updatedBooking, nil
}

func (s bookingService) DeleteBooking(ctx context.Context, bookingId string) error {
	uid, err := uuid.Parse(bookingId)
	if err != nil {
		return fmt.Errorf("DeleteBooking: invalid booking id: %w", err)
	}
	if err := s.repo.DeleteBookingById(uid); err != nil {
		return fmt.Errorf("DeleteBooking: %w", err)
	}
	return nil
}

func (s bookingService) GetAllBookingsByUser(ctx context.Context, userId string) ([]data.Booking, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("GetAllBookingsByUser: invalid user id: %w", err)
	}
	bookings, err := s.repo.GetAllBookingsByUser(uid)
	if err != nil {
		return nil, fmt.Errorf("GetAllBookingsByUser: %w", err)
	}

	return bookings, nil
}

func (s bookingService) GetAllBookings(ctx context.Context) ([]data.Booking, error) {
	bookings, err := s.repo.GetAllBookings()
	if err != nil {
		return nil, fmt.Errorf("GetAllBookings: %w", err)
	}

	return bookings, nil
}

func NewBookingService(repo data.BookingRepository, conv converter.Converter) BookingService {
	return &bookingService{
		repo: repo,
		conv: conv,
	}
}
