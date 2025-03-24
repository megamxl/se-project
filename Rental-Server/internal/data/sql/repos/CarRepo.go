package repos

import (
	"context"
	"errors"
	dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/model"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/query"
	"time"
)

var _ dataInt.CarRepository = (*CarRepo)(nil)

type CarRepo struct {
	Q   *query.Query
	Ctx context.Context
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
	_, err := c.Q.WithContext(c.Ctx).Car.Where(c.Q.Car.Vin.Eq(car.Vin)).Updates(car)
	if err != nil {
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
	//TODO implement me
	// I need the superiors (Maxls) help to autogenerate the relation
	panic("implement me")
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
