package repos

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/model"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/query"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"
)

var _ dataInt.CarRepository = (*CarRepo)(nil)

type CarRepo struct {
	Q   *query.Query
	Ctx context.Context
	Db  *gorm.DB
}

func (c CarRepo) GetCarsNotInList(vins []string) ([]dataInt.Car, error) {

	slog.Info("the vins len ", vins[0])

	find, err := c.Q.WithContext(c.Ctx).Car.Where(c.Q.Car.Vin.NotIn(vins...)).Find()
	if err != nil {
		return nil, err
	}

	var cars []dataInt.Car

	for _, v := range find {
		cars = append(cars, modelToIntCar(v))
	}

	return cars, nil
}

func (c CarRepo) GetCarByVin(vin string) (dataInt.Car, error) {
	find, err := c.Q.WithContext(c.Ctx).Car.Where(c.Q.Car.Vin.Eq(vin)).Find()
	if err != nil {
		return dataInt.Car{}, err
	}

	if len(find) != 1 {
		return dataInt.Car{}, errors.New("Car not found or duplicate VIN")
	}

	return modelToIntCar(find[0]), nil
}

func (c CarRepo) SaveCar(car dataInt.Car) (dataInt.Car, error) {
	toSaveCar := intToModelCar(car)
	err := c.Q.WithContext(c.Ctx).Car.Save(toSaveCar)
	if err != nil {
		return dataInt.Car{}, err
	}
	return modelToIntCar(toSaveCar), nil
}

func (c CarRepo) UpdateCar(car dataInt.Car) (dataInt.Car, error) {
	status, err := c.Q.WithContext(c.Ctx).Car.Where(c.Q.Car.Vin.Eq(car.Vin)).Updates(car)
	if err != nil || status.RowsAffected == 0 {
		if status.RowsAffected == 0 {
			err = errors.New("car not found")
		}
		return dataInt.Car{}, err
	}

	find, err := c.Q.WithContext(c.Ctx).Car.Where(c.Q.Car.Vin.Eq(car.Vin)).Find()
	return modelToIntCar(find[0]), nil
}

func (c CarRepo) DeleteCarByVin(vin string) error {
	_, err := c.Q.WithContext(c.Ctx).Car.Where(c.Q.Car.Vin.Eq(vin)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (c CarRepo) GetCarsAvailableInTimeRange(startTime time.Time, endTime time.Time) ([]dataInt.Car, error) {
	// I need the superiors (Maxls) help to autogenerate the relation
	layout := "2006-01-02 15:04:05"
	var cars []dataInt.Car

	if os.Getenv("BOOKING_SERVICE_URL") != "" {
		booked, err := getVinsBooked(startTime, endTime)

		if err != nil {
			return nil, err
		}

		list, err := c.GetCarsNotInList(booked)

		if err != nil {
			return nil, err
		}

		return list, nil

	} else {
		tx := c.Db.WithContext(c.Ctx).Raw("SELECT c.* FROM car c WHERE NOT EXISTS ( SELECT 1 FROM booking b WHERE b.car_vin = c.vin AND tsrange(b.start_time, b.end_time) && tsrange(@start, @end))", sql.Named("start", startTime.Format(layout)), sql.Named("end", endTime.Format(layout))).Find(&cars)

		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return cars, nil
}

func getVinsBooked(startTime, endTime time.Time) ([]string, error) {
	endpoint := os.Getenv("BOOKING_SERVICE_URL") + "/bookings/rpc/in_range"

	// Construct the query parameters
	params := url.Values{}

	params.Add("startTime", startTime.Format("2006-01-02"))
	params.Add("endTime", endTime.Format("2006-01-02"))

	fullURL := endpoint
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	// Make the HTTP GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	type Booking struct {
		VIN        *string  `json:"VIN,omitempty"`
		BookingId  *string  `json:"bookingId,omitempty"`
		Currency   *string  `json:"currency,omitempty"`
		PaidAmount *float32 `json:"paidAmount,omitempty"`
		Status     *string  `json:"status,omitempty"`
		UserId     *string  `json:"userId,omitempty"`
	}

	// Decode the JSON response
	var bookingList []Booking
	err = json.NewDecoder(resp.Body).Decode(&bookingList)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	slog.Info("booking list: ", bookingList)

	var result []string
	for _, b := range bookingList {
		result = append(result, *b.VIN)
	}

	return result, nil
}

func intToModelCar(car dataInt.Car) *model.Car {
	newCar := &model.Car{
		Vin:         car.Vin,
		Model:       car.Model,
		Brand:       car.Brand,
		ImageURL:    car.ImageUrl,
		PricePerDay: car.PricePerDay,
	}
	return newCar
}

func modelToIntCar(newCar *model.Car) dataInt.Car {
	savedCar := dataInt.Car{
		Vin:         newCar.Vin,
		Model:       newCar.Model,
		Brand:       newCar.Brand,
		ImageUrl:    newCar.ImageURL,
		PricePerDay: newCar.PricePerDay,
	}
	return savedCar
}
