package carEvents

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"log/slog"
)

type PulsarProducer struct {
	Producer pulsar.Producer
}

func (i PulsarProducer) AddCar(car data.Car) {
	sendMessage(car, i.Producer, "ADD")
}

func sendMessage(car data.Car, producer pulsar.Producer, op string) {
	msg, err := generateEventJsonAsBytes(car, op)
	if err != nil {
		return
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: msg,
	})

	if err != nil {
		slog.Error("Failed to send message", err)
		return
	}
}

func (i PulsarProducer) RemoveCar(car data.Car) {
	sendMessage(car, i.Producer, "REMOVE")
}

func (i PulsarProducer) UpdateCar(car data.Car) {
	sendMessage(car, i.Producer, "UPDATE")
}

func generateEventJsonAsBytes(car data.Car, op string) ([]byte, error) {
	eventStruct := CarEventStruct{
		Car:       car,
		Operation: op,
	}

	bytes, err := json.Marshal(eventStruct)
	if err != nil {
		slog.Error(fmt.Sprintf("Error marshalling car event struct for operation: %s", op))
		return nil, err
	}
	return bytes, nil
}

var _ CarEvent = (*PulsarProducer)(nil)
