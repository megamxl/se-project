package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/megamxl/se-project/Rental-Server/api/DTO"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
)

type BookingService interface {
	BookCar(ctx context.Context, req DTO.Booking) (DTO.Booking, error)
	GetBookingById(ctx context.Context, id string) (DTO.Booking, error)
	UpdateBooking(ctx context.Context, bookingId string, status string) (DTO.Booking, error)
	DeleteBooking(ctx context.Context, bookingId string) error
	GetAllBookingsByUser(ctx context.Context, userId string) ([]DTO.Booking, error)
	GetAllBookings(ctx context.Context) ([]DTO.Booking, error)
}

type bookingService struct {
	repo data.BookingRepository
}

func (s bookingService) BookCar(ctx context.Context, req DTO.Booking) (DTO.Booking, error) {
	if req.CarVin == "" {
		return DTO.Booking{}, errors.New("BookCar: VIN is empty")
	}

	dataBooking := MapDTOBookingToDataBooking(req)

	if dataBooking.Status == "" {
		dataBooking.Status = "pending"
	}

	saved, err := s.repo.SaveBooking(dataBooking)
	if err != nil {
		return DTO.Booking{}, fmt.Errorf("BookCar: %w", err)
	}

	return MapDataBookingToDTOBooking(saved), nil
}

func (s bookingService) GetBookingById(ctx context.Context, id string) (DTO.Booking, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return DTO.Booking{}, fmt.Errorf("GetBookingById: invalid id format: %w", err)
	}
	booking, err := s.repo.GetBookingById(uid)
	if err != nil {
		return DTO.Booking{}, fmt.Errorf("GetBookingById: %w", err)
	}
	return MapDataBookingToDTOBooking(booking), nil
}

func (s bookingService) UpdateBooking(ctx context.Context, bookingId string, status string) (DTO.Booking, error) {
	uid, err := uuid.Parse(bookingId)
	if err != nil {
		return DTO.Booking{}, fmt.Errorf("UpdateBooking: invalid booking id: %w", err)
	}

	booking, err := s.repo.GetBookingById(uid)
	if err != nil {
		return DTO.Booking{}, fmt.Errorf("UpdateBooking: %w", err)
	}
	booking.Status = status

	updated, err := s.repo.SaveBooking(booking)
	if err != nil {
		return DTO.Booking{}, fmt.Errorf("UpdateBooking: %w", err)
	}
	return MapDataBookingToDTOBooking(updated), nil
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

func (s bookingService) GetAllBookingsByUser(ctx context.Context, userId string) ([]DTO.Booking, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("GetAllBookingsByUser: invalid user id: %w", err)
	}
	bookings, err := s.repo.GetAllBookingsByUser(uid)
	if err != nil {
		return nil, fmt.Errorf("GetAllBookingsByUser: %w", err)
	}
	dtoBookings := make([]DTO.Booking, len(bookings))
	for i, b := range bookings {
		dtoBookings[i] = MapDataBookingToDTOBooking(b)
	}
	return dtoBookings, nil
}

func (s bookingService) GetAllBookings(ctx context.Context) ([]DTO.Booking, error) {
	bookings, err := s.repo.GetAllBookings()
	if err != nil {
		return nil, fmt.Errorf("GetAllBookings: %w", err)
	}
	dtoBookings := make([]DTO.Booking, len(bookings))
	for i, b := range bookings {
		dtoBookings[i] = MapDataBookingToDTOBooking(b)
	}
	return dtoBookings, nil
}

func NewBookingService(repo data.BookingRepository) BookingService {
	return &bookingService{
		repo: repo,
	}
}

func MapDTOBookingToDataBooking(dto DTO.Booking) data.Booking {
	return data.Booking{
		Id:        dto.Id,
		CarVin:    dto.CarVin,
		UserId:    dto.UserId,
		StartTime: dto.StartTime,
		EndTime:   dto.EndTime,
		Status:    dto.Status,
	}
}

func MapDataBookingToDTOBooking(dto data.Booking) DTO.Booking {
	return DTO.Booking{
		Id:        dto.Id,
		CarVin:    dto.CarVin,
		UserId:    dto.UserId,
		StartTime: dto.StartTime,
		EndTime:   dto.EndTime,
		Status:    dto.Status,
	}
}
