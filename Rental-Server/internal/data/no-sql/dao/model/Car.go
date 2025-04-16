package model

type Car struct {
	Vin         string  `bson:"_id" json:"vin"`
	Model       string  `bson:"model" json:"model"`
	Brand       string  `bson:"brand" json:"brand"`
	ImageUrl    string  `bson:"image_url" json:"image_url"`
	PricePerDay float64 `bson:"price_per_day" json:"price_per_day"`
}
