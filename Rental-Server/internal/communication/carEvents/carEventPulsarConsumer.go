package carEvents

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"log/slog"
)

type pulsarConsumer struct {
	repo     data.CarRepository
	reader   pulsar.Reader
	consumer pulsar.Consumer
}

func (i pulsarConsumer) AddCar(car data.Car) {

	update, err := i.repo.SaveCar(car)

	if err != nil {
		slog.Error("Error while saving car in Pulsar listener ", err)
		return
	}

	slog.Info(fmt.Sprintf("Updated car in Pulsar listener %s", update.Vin))

}

func (i pulsarConsumer) RemoveCar(car data.Car) {

	err := i.repo.DeleteCarByVin(car.Vin)

	if err != nil {
		slog.Error("Error while deleting car in Pulsar listener ", err)
		return
	}

	slog.Info(fmt.Sprintf("Updated car in Pulsar listener %s", car.Vin))

}

func (i pulsarConsumer) UpdateCar(car data.Car) {
	update, err := i.repo.UpdateCar(car)

	if err != nil {
		slog.Error("Error while Updating car in Pulsar listener ", err)
		return
	}

	slog.Info(fmt.Sprintf("Updated car in Pulsar listener %s", update.Vin))
}

func NewPulsarConsumer(reader pulsar.Reader, repo data.CarRepository) CarEvent {
	consumer := &pulsarConsumer{
		repo:   repo,
		reader: reader,
	}

	go func() {
		slog.Debug("Waiting for message...")
		for {
			msg, err := reader.Next(context.Background())
			if err != nil {
				slog.Error("Error reading from Pulsar", err)
				continue
			}

			var car CarEventStruct
			if err := json.Unmarshal(msg.Payload(), &car); err != nil {
				slog.Error("Failed to unmarshal car event", err)
				continue
			}

			switch car.Operation {
			case "ADD":
				consumer.AddCar(car.Car)
			case "REMOVE":
				consumer.RemoveCar(car.Car)
			case "UPDATE":
				consumer.UpdateCar(car.Car)
			default:
				slog.Warn("Unknown eventType", slog.String("eventType", car.Operation))
			}
		}
	}()

	return consumer
}

var _ CarEvent = (*pulsarConsumer)(nil)
