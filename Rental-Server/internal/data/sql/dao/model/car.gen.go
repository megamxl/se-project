// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCar = "car"

// Car mapped from table <car>
type Car struct {
	Vin         string    `gorm:"column:vin;primaryKey" json:"vin"`
	Model       string    `gorm:"column:model;not null" json:"model"`
	Brand       string    `gorm:"column:brand;not null" json:"brand"`
	ImageURL    string    `gorm:"column:image_url" json:"image_url"`
	PricePerDay float64   `gorm:"column:price_per_day;not null" json:"price_per_day"`
	CreatedAt   time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
}

// TableName Car's table name
func (*Car) TableName() string {
	return TableNameCar
}
