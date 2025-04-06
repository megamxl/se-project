package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
)

type BookingService interface {
	BookCar(ctx context.Context, req data.Booking) (data.Booking, error)
	GetBookingById(ctx context.Context, id string) (data.Booking, error)
	UpdateBooking(ctx context.Context, bookingId string, status string) (data.Booking, error)
	DeleteBooking(ctx context.Context, bookingId string) error
	GetAllBookingsByUser(ctx context.Context, userId string) ([]data.Booking, error)
	GetAllBookings(ctx context.Context) ([]data.Booking, error)
}

type bookingService struct {
	repo data.BookingRepository
}

func (s bookingService) BookCar(ctx context.Context, req data.Booking) (data.Booking, error) {
	if req.CarVin == "" {
		return data.Booking{}, errors.New("BookCar: VIN is empty")
	}

	if req.Status == "" {
		req.Status = "pending"
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

	updatedBooking, err := s.repo.UpdateBookingStateById(booking.Id, status)
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

func NewBookingService(repo data.BookingRepository) BookingService {
	return &bookingService{
		repo: repo,
	}
}
