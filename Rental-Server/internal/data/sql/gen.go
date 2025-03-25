package sql

import (
	"context"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/query"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/repos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"time"
)

//go:generate  gentool -c "./gen.tool"

func Db() {
	// Connect to your database

	dsn := "host=localhost user=admin password=admin dbname=main port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	use := query.Use(db)
	//repo := repos.RentalRepo{
	//	Q:   use,
	//	Ctx: context.Background(),
	//}
	//
	//id, err := repo.GetBookingById(uuid.MustParse("49b68d7f-069b-44b3-8863-00e69d3be9f7"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//booking := dataInt.Booking{
	//	CarVin:    "1HGBH41JXMN109186",
	//	UserId:    uuid.MustParse("b9d76281-8850-46dd-b6ce-bc59259a52ab"),
	//	StartTime: time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC),
	//	EndTime:   time.Date(2025, 5, 27, 0, 0, 0, 0, time.UTC),
	//	Status:    "confirmed",
	//}

	carRepo := repos.CarRepo{
		Db:  db,
		Ctx: context.Background(),
		Q:   use,
	}

	startTime := time.Date(2025, 3, 25, 10, 0, 0, 0, time.UTC)
	endTime := time.Date(2025, 3, 25, 14, 0, 0, 0, time.UTC)

	_, err = carRepo.GetCarsAvailableInTimeRange(startTime, endTime)
	if err != nil {
		slog.Error(err.Error())
	}

}
