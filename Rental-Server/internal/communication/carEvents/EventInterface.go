package carEvents

import dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"

type CarEvent interface {
	AddCar(car dataInt.Car)
	RemoveCar(car dataInt.Car)
	UpdateCar(car dataInt.Car)
}

type CarEventStruct struct {
	Car       dataInt.Car
	Operation string
}
