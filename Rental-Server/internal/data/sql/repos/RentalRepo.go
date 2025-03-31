package repos

import (
	"context"
	"errors"
	"github.com/google/uuid"
	dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/model"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/query"
)

type RentalRepo struct {
	Q   *query.Query
	Ctx context.Context
}

func (r RentalRepo) UpdateBookingStateById(id uuid.UUID, state string) (dataInt.Booking, error) {

	_, err := r.Q.WithContext(r.Ctx).Booking.Where(r.Q.Booking.ID.Eq(id.String())).Update(
		r.Q.Booking.Status, state)
	if err != nil {
		return dataInt.Booking{}, err
	}

	first, err := r.Q.WithContext(r.Ctx).Booking.
		Where(r.Q.Booking.ID.Eq(id.String())).
		First()
	if err != nil {
		return dataInt.Booking{}, err
	}

	return modelToInt(first), nil

}

func (r RentalRepo) GetAllBookingsByUser(userId uuid.UUID) ([]dataInt.Booking, error) {

	find, err := r.Q.WithContext(r.Ctx).Booking.Where(r.Q.Booking.CustomerID.Eq(userId.String())).Find()
	if err != nil {
		return nil, err
	}

	return MapBookingsToDTOs(find), nil

}

func (r RentalRepo) GetAllBookings() ([]dataInt.Booking, error) {

	find, err := r.Q.WithContext(r.Ctx).Booking.Find()
	if err != nil {
		return nil, err
	}

	return MapBookingsToDTOs(find), nil
}

func (r RentalRepo) GetBookingsByVin(vin string) ([]dataInt.Booking, error) {

	find, err := r.Q.WithContext(r.Ctx).Booking.Where(r.Q.Booking.CarVin.Eq(vin)).Find()
	if err != nil {
		return nil, err
	}

	return MapBookingsToDTOs(find), nil

}

func (r RentalRepo) GetBookingById(id uuid.UUID) (dataInt.Booking, error) {

	find, err := r.Q.WithContext(r.Ctx).Booking.Where(r.Q.Booking.ID.Eq(id.String())).Find()
	if err != nil {
		return dataInt.Booking{}, err
	}

	if len(find) != 1 {
		return dataInt.Booking{}, errors.New("booking not found or duplicate ID")
	}

	return modelToInt(find[0]), nil
}

func (r RentalRepo) SaveBooking(booking dataInt.Booking) (dataInt.Booking, error) {

	newBooking := intToModel(booking)
	newBooking.ID = uuid.NewString()

	err := r.Q.WithContext(r.Ctx).Booking.Save(newBooking)
	if err != nil {
		return dataInt.Booking{}, err
	}

	savedBooking := modelToInt(newBooking)

	return savedBooking, nil

}

func (r RentalRepo) DeleteBookingsByVin(vin string) error {

	_, err := r.Q.WithContext(r.Ctx).Booking.Where(r.Q.Booking.CarVin.Eq(vin)).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (r RentalRepo) DeleteBookingById(id uuid.UUID) error {
	_, err := r.Q.WithContext(r.Ctx).Booking.Where(r.Q.Booking.ID.Eq(id.String())).Delete()
	if err != nil {
		return err
	}
	return nil
}

var _ dataInt.BookingRepository = (*RentalRepo)(nil)

func intToModel(booking dataInt.Booking) *model.Booking {
	newBooking := &model.Booking{
		CarVin:     booking.CarVin,
		CustomerID: booking.UserId.String(),
		StartTime:  booking.StartTime,
		EndTime:    booking.EndTime,
		Status:     booking.Status,
		Paidamount: booking.AmountPaid,
		Currency:   booking.Currency,
	}
	return newBooking
}

func modelToInt(newBooking *model.Booking) dataInt.Booking {
	savedBooking := dataInt.Booking{
		Id:         uuid.MustParse(newBooking.ID),
		CarVin:     newBooking.CarVin,
		UserId:     uuid.MustParse(newBooking.CustomerID),
		StartTime:  newBooking.StartTime,
		EndTime:    newBooking.EndTime,
		Status:     newBooking.Status,
		AmountPaid: newBooking.Paidamount,
		Currency:   newBooking.Currency,
	}
	return savedBooking
}

func MapBookingsToDTOs(bookings []*model.Booking) []dataInt.Booking {
	dtos := make([]dataInt.Booking, len(bookings))
	for i, booking := range bookings {
		dtos[i] = modelToInt(booking)
	}
	return dtos
}
