package carEvents

import (
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/service"
)

type pulsarConsumer struct {
	service service.CarService
}

func (i pulsarConsumer) AddCar(car data.Car) {
	//TODO implement me
	panic("implement me")
}

func (i pulsarConsumer) RemoveCar(car data.Car) {
	//TODO implement me
	panic("implement me")
}

func (i pulsarConsumer) UpdateCar(car data.Car) {
	//TODO implement me
	panic("implement me")
}

var _ CarEvent = (*pulsarConsumer)(nil)
