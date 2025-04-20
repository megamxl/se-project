package di

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/carEvents"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/converter"
	myGrpcImpl "github.com/megamxl/se-project/Rental-Server/internal/communication/converter/grpc"
	myGrpcStub "github.com/megamxl/se-project/Rental-Server/internal/communication/converter/grpc/proto"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/converter/soap"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/query"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/repos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"
)

func GetSQLDatabaseConnection() *gorm.DB {

	postgresEnv := os.Getenv("POSTGRES_DNS")

	if postgresEnv == "" {
		return nil
	}

	db, err := gorm.Open(postgres.Open(postgresEnv), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func SeedDatabase(db *gorm.DB) {

	seeder := sql.NewSqlSeeder(db)

	monoSeeder, err := strconv.ParseBool(os.Getenv("SEED_MONOLITH"))
	if err != nil {
		slog.Info("Not seeding database for Monolith since Argument can't be parsed as boolean")
	}

	if monoSeeder {
		seeder.SeedUser()
		seeder.SeedCars()
		seeder.SeedBooking()
		seeder.SeedBookingMonolithConstraints()
		slog.Info("✅ seeded the tables for Monolith")
		return
	}

	userSql, err := strconv.ParseBool(os.Getenv("SEED_USER_SQL"))
	if err != nil {
		slog.Info("Not seeding database for SEED_USER_SQL since Argument can't be parsed as boolean")
	}

	if userSql {
		seeder.SeedUser()
		slog.Info("✅ seeded the user Table")
		return
	}

	carSQL, err := strconv.ParseBool(os.Getenv("SEED_CAR_SQL"))
	if err != nil {
		slog.Info("Not seeding database for SEED_CAR_SQL since Argument can't be parsed as boolean")
	}

	if carSQL {
		seeder.SeedCars()
		slog.Info("✅ seeded the Car Table")
	}

	bookSQL, err := strconv.ParseBool(os.Getenv("SEED_BOOKING_SQL"))
	if err != nil {
		slog.Info("Not seeding database for SEED_BOOKING_SQL since Argument can't be parsed as boolean")
	}

	if bookSQL {
		seeder.SeedBooking()
		seeder.SeedBookingConstraints()
		slog.Info("✅ seeded the Booking Table")
	}
}

func GetUserRepositorySQL(db *gorm.DB) data.UserRepository {
	return &repos.UserRepo{
		Q:   query.Use(db),
		Ctx: context.Background(),
	}
}

func GetBookingRepositorySQL(db *gorm.DB) data.BookingRepository {
	return repos.RentalRepo{
		Q:   query.Use(db),
		Ctx: context.Background(),
	}
}

func GetCarRepositorySQL(db *gorm.DB) data.CarRepository {
	return &repos.CarRepo{
		Db:  db,
		Ctx: context.Background(),
		Q:   query.Use(db),
	}
}

func GetConvertor() converter.Converter {
	var convService converter.Converter

	if os.Getenv("CONVERTOR_SOAP_URL") != "" {
		convService = soap.NewSoapService(os.Getenv("CONVERTOR_SOAP_URL"))

		slog.Info("Connected to SOAP server and using it ")

	} else if os.Getenv("CONVERTOR_GRPC_URL") != "" {

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.NewClient(os.Getenv("CONVERTOR_GRPC_URL"), opts...)
		if err != nil {
			log.Fatal("Failed to connect to server:", err)
		}

		client := myGrpcStub.NewConvertorClient(conn)

		convService = myGrpcImpl.NewConverter(client)
		slog.Info("Connected to GRPC server and using it ")
	} else {
		log.Fatal("Unsupported config no convertor set. Set CONVERTOR_SOAP_URL or CONVERTOR_GRPC_URL if you want an convertor")
	}

	return convService
}

func PulsarListner(repository data.CarRepository) error {

	if repository == nil {
		log.Fatal("CarRepository is nil when setting up Pulsar")
	}

	if os.Getenv("PULSAR_LISTENER") == "true" {

		client, err := pulsar.NewClient(pulsar.ClientOptions{
			URL:               os.Getenv("PULSAR_URL"),
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
		if err != nil {
			slog.Error("Could not instantiate Pulsar client: %v", err)
			return err
		}

		reader, err := client.CreateReader(pulsar.ReaderOptions{
			Topic:          "car-events",
			StartMessageID: pulsar.EarliestMessageID(),
		})

		if err != nil {
			return err
		}

		carEvents.NewPulsarConsumer(reader, repository)
	}
	return nil
}
