package sql

import (
	"gorm.io/gorm"
	"log"
	"os"
)

type Seeder struct {
	DB *gorm.DB
}

func (s Seeder) SeedUser() {
	runSQLFile(s, "internal/data/sql/user.sql")
}
func (s Seeder) SeedCars() {
	runSQLFile(s, "internal/data/sql/car.sql")
}
func (s Seeder) SeedBooking() {
	runSQLFile(s, "internal/data/sql/booking.sql")
}

func (s Seeder) SeedBookingMonolithConstraints() {
	runSQLFile(s, "internal/data/sql/monolith-booking-constraints.sql")
}

func (s Seeder) SeedBookingConstraints() {
	runSQLFile(s, "internal/data/sql/booking_constraints.sql")
}

func runSQLFile(s Seeder, fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Cant Open "+fileName, err)
	}

	if err := s.DB.Exec(string(file)).Error; err != nil {
		panic("failed to execute " + fileName + " SQL: " + err.Error())
	}
}

func NewSqlSeeder(db *gorm.DB) Seeder {
	return Seeder{
		DB: db,
	}
}
