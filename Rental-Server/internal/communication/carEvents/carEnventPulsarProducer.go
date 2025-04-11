package carEvents

import "github.com/megamxl/se-project/Rental-Server/internal/data"

type pulsarProducer struct{}

func (i pulsarProducer) AddCar(car data.Car) {
	//TODO implement me
	panic("implement me")
}

func (i pulsarProducer) RemoveCar(car data.Car) {
	//TODO implement me
	panic("implement me")
}

func (i pulsarProducer) UpdateCar(car data.Car) {
	//TODO implement me
	panic("implement me")
}

var _ CarEvent = (*pulsarProducer)(nil)
