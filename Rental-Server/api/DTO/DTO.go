package DTO

import (
	"github.com/google/uuid"
	"time"
)

type RentalUser struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type Car struct {
	Vin         string `json:"vin"`
	Model       string `json:"model"`
	Brand       string `json:"brand"`
	ImageUrl    string `json:"imageUrl"`
	PricePerDay string `json:"pricePerDay"`
}

type Booking struct {
	Id        uuid.UUID `json:"id"`
	CarVin    string    `json:"carVin"`
	UserId    uuid.UUID `json:"userId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	//look into enums
	Status string `json:"status"`
}
