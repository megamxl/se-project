package data

import (
	"github.com/google/uuid"
	"time"
)

type RentalUser struct {
	Id       uuid.UUID
	Name     string
	Email    string
	Password string
	Admin    bool
}

type Car struct {
	Vin         string
	Model       string
	Brand       string
	ImageUrl    string
	PricePerDay float64
}

type Booking struct {
	Id        uuid.UUID
	CarVin    string
	UserId    uuid.UUID
	StartTime time.Time
	EndTime   time.Time
	//look into enums
	Status     string
	AmountPaid float64
	Currency   string
}

type UserRepository interface {
	GetUserByEmail(email string) (RentalUser, error)
	GetUserById(id uuid.UUID) (RentalUser, error)
	GetAllUsers() ([]RentalUser, error)
	UpdateUserById(id uuid.UUID, update RentalUser) (RentalUser, error)
	UpdateUserByEmail(email string, update RentalUser) (RentalUser, error)
	DeleteUserById(id uuid.UUID) error
	SaveUser(user RentalUser) (RentalUser, error)
}

type CarRepository interface {
	GetCarByVin(vin string) (Car, error)
	SaveCar(car Car) (Car, error)
	UpdateCar(car Car) (Car, error)
	DeleteCarByVin(vin string) error

	GetCarsAvailableInTimeRange(startTime time.Time, endTime time.Time) ([]Car, error)
	GetCarsNotInList(vins []string) ([]Car, error)
	GetAllCars() ([]Car, error)
}

type BookingRepository interface {
	GetBookingsByVin(vin string) ([]Booking, error)
	GetBookingById(id uuid.UUID) (Booking, error)
	SaveBooking(booking Booking) (Booking, error)
	DeleteBookingsByVin(vin string) error
	DeleteBookingById(id uuid.UUID) error
	GetAllBookingsByUser(userId uuid.UUID) ([]Booking, error)
	GetAllBookings() ([]Booking, error)
	UpdateBookingStateById(id uuid.UUID, update string) (Booking, error)
	GetBookingsInTimeRange(startTime time.Time, endTime time.Time) ([]Booking, error)
}
