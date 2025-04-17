package repos

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/no-sql/dao/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ dataInt.CarRepository = (*CarRepo)(nil)

type CarRepo struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func (c CarRepo) GetCarsNotInList(vins []string) ([]dataInt.Car, error) {

	filter := bson.M{}

	if len(vins) > 0 {
		filter["_id"] = bson.M{
			"$nin": vins,
		}
	}

	cursor, err := c.Collection.Find(c.Ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.Ctx)

	var cars []model.Car
	if err := cursor.All(c.Ctx, &cars); err != nil {
		return nil, err
	}

	var result []dataInt.Car
	for _, car := range cars {
		result = append(result, convertModelToDataCar(car))
	}

	return result, nil
}

func (c CarRepo) GetCarByVin(vin string) (dataInt.Car, error) {
	log.Printf("🚗 [Mongo] GetCarByVin: %s", vin)
	var result model.Car
	filter := bson.M{"_id": vin}
	err := c.Collection.FindOne(c.Ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dataInt.Car{}, errors.New("car not found")
		}
		return dataInt.Car{}, err
	}
	return convertModelToDataCar(result), nil
}

func (c CarRepo) SaveCar(car dataInt.Car) (dataInt.Car, error) {
	log.Printf("🚗 [Mongo] SaveCar: %+v", car)
	if car.Vin == "" {
		return dataInt.Car{}, errors.New("VIN must be set")
	}

	doc := model.Car{
		Vin:         car.Vin,
		Model:       car.Model,
		Brand:       car.Brand,
		ImageUrl:    car.ImageUrl,
		PricePerDay: car.PricePerDay,
	}

	_, err := c.Collection.InsertOne(c.Ctx, doc)
	if err != nil {
		return dataInt.Car{}, err
	}

	return car, nil
}

func (c CarRepo) UpdateCar(car dataInt.Car) (dataInt.Car, error) {
	filter := bson.M{"_id": car.Vin}
	update := bson.M{"$set": bson.M{
		"model":         car.Model,
		"brand":         car.Brand,
		"image_url":     car.ImageUrl,
		"price_per_day": car.PricePerDay,
	}}

	_, err := c.Collection.UpdateOne(c.Ctx, filter, update)
	if err != nil {
		return dataInt.Car{}, err
	}

	log.Printf("🚗 [Mongo] UpdateCar: %+v", car)

	return car, nil
}

func (c CarRepo) DeleteCarByVin(vin string) error {
	log.Printf("🚗 [Mongo] DeleteCarByVin: %s", vin)
	filter := bson.M{"_id": vin}
	_, err := c.Collection.DeleteOne(c.Ctx, filter)
	return err
}

func (c CarRepo) GetCarsAvailableInTimeRange(startTime time.Time, endTime time.Time) ([]dataInt.Car, error) {
	// This is a placeholder. Real availability logic involves checking bookings.
	if os.Getenv("BOOKING_SERVICE_URL") == "" {
		return nil, errors.New("BOOKING_SERVICE_URL not set cant get bookings")
	}

	log.Printf("🚗 [Mongo] GetCarsAvailableInTimeRange: from %s to %s", startTime.Format("2006-01-02"), endTime.Format("2006-01-02"))
	booked, err := dataInt.GetVinsBooked(startTime, endTime)

	if err != nil {
		return nil, err
	}

	list, err := c.GetCarsNotInList(booked)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Helper

func NewCarRepo(ctx context.Context, db *mongo.Database) *CarRepo {
	log.Println("📦 [Mongo] NewCarRepo")
	return &CarRepo{
		Collection: db.Collection("cars"),
		Ctx:        ctx,
	}
}

func convertModelToDataCar(m model.Car) dataInt.Car {
	return dataInt.Car{
		Vin:         m.Vin,
		Model:       m.Model,
		Brand:       m.Brand,
		ImageUrl:    m.ImageUrl,
		PricePerDay: m.PricePerDay,
	}
}
