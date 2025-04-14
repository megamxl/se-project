package api

import (
	"context"
	"errors"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/google/uuid"
	"github.com/megamxl/se-project/Rental-Server/api/Util"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/carEvents"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/converter"
	myGrpcImpl "github.com/megamxl/se-project/Rental-Server/internal/communication/converter/grpc"
	myGrpcStub "github.com/megamxl/se-project/Rental-Server/internal/communication/converter/grpc/proto"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/megamxl/se-project/Rental-Server/internal/communication/converter/soap"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/query"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/repos"
	"github.com/megamxl/se-project/Rental-Server/internal/middleware"
	"github.com/megamxl/se-project/Rental-Server/internal/service"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	nosqldb "github.com/megamxl/se-project/Rental-Server/internal/data/no-sql/db"
	nosqlrepos "github.com/megamxl/se-project/Rental-Server/internal/data/no-sql/repos"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct {
	carService     service.CarService
	userService    service.UserService
	bookingService service.BookingService
}

func (s Server) Login(w http.ResponseWriter, r *http.Request) {
	var body LoginJSONRequestBody

	if err := Util.DecodeJSON(r, &body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	errMsg := "Password or Email incorrect"

	user, err := s.userService.GetUserByEmail(r.Context(), string(body.Email))
	if err != nil {
		http.Error(w, errMsg, http.StatusUnauthorized)
	}

	if user.Password != body.Password {
		http.Error(w, errMsg, http.StatusUnauthorized)
	}

	jwtForUser, err := middleware.CreateJWForUser(user)
	if err != nil {
		http.Error(w, "can't create Token check wit the support", http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   jwtForUser,
		Expires: time.Now().Add(24 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
}

func (s Server) DeleteBooking(w http.ResponseWriter, r *http.Request, params DeleteBookingParams) {
	if params.BookingId == "" {
		http.Error(w, "missing 'bookingId' query parameter", http.StatusBadRequest)
		return
	}

	if err := s.bookingService.DeleteBooking(r.Context(), params.BookingId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s Server) GetBookings(w http.ResponseWriter, r *http.Request) {

	userIdFromRequest, err := getUserIdFromRequest(r)
	if err != nil {
		http.Error(w, "User not found contact support", http.StatusBadRequest)
	}

	dataBookings, err := s.bookingService.GetAllBookingsByUser(r.Context(), userIdFromRequest.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookings := make([]Booking, len(dataBookings))
	for index, db := range dataBookings {
		bookings[index] = MapDataBookingToBooking(db)
	}

	if err := Util.WriteJSON(w, http.StatusOK, bookings); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func (s Server) BookCar(w http.ResponseWriter, r *http.Request) {
	var req BookCarJSONBody
	if err := Util.DecodeJSON(r, &req); err != nil {
		http.Error(w, "invalid Request body", http.StatusBadRequest)
		return
	}

	userIdFromRequest, err := getUserIdFromRequest(r)
	if err != nil {
		http.Error(w, "User not found contact support", http.StatusBadRequest)
	}

	vin, err := s.carService.GetCarByVin(r.Context(), *req.VIN)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	booking, err := s.bookingService.BookCar(r.Context(), data.Booking{
		CarVin:    *req.VIN,
		UserId:    userIdFromRequest,
		StartTime: req.StartTime.Time,
		EndTime:   req.EndTime.Time,
	},
		string(*req.Currency),
		vin.PricePerDay,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := Util.WriteJSON(w, http.StatusOK, MapDataBookingToBooking(booking)); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s Server) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BookingId string `json:"bookingId"`
		Status    string `json:"status"`
	}
	if err := Util.DecodeJSON(r, &req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.BookingId == "" || req.Status == "" {
		http.Error(w, "missing bookingId or status", http.StatusBadRequest)
		return
	}

	_, err := s.bookingService.UpdateBooking(r.Context(), req.BookingId, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) GetBookingById(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "missing id in path", http.StatusBadRequest)
		return
	}

	booking, err := s.bookingService.GetBookingById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := Util.WriteJSON(w, http.StatusOK, MapDataBookingToBooking(booking)); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func (s Server) GetAllBookingsByUser(w http.ResponseWriter, r *http.Request) {
	dataBookings, err := s.bookingService.GetAllBookings(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookings := make([]Booking, len(dataBookings))
	for index, db := range dataBookings {
		bookings[index] = MapDataBookingToBooking(db)
	}

	if err := Util.WriteJSON(w, http.StatusOK, bookings); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusBadRequest)
		return
	}
}

func (s Server) DeleteCar(w http.ResponseWriter, r *http.Request, params DeleteCarParams) {
	if params.VIN == "" {
		http.Error(w, "ERROR: VIN parameter is required", http.StatusBadRequest)
		return
	}

	if err := s.carService.DeleteCarByVin(r.Context(), params.VIN); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s Server) ListCars(w http.ResponseWriter, r *http.Request, params ListCarsParams) {
	if params.Currency == "" {
		http.Error(w, "ERROR: currency parameter is required", http.StatusBadRequest)
		return
	}

	var startTimeVal time.Time
	var endTimeVal time.Time

	if params.StartTime != nil {
		startTimeVal = params.StartTime.Time
	}

	if params.EndTime != nil {
		endTimeVal = params.EndTime.Time
	}

	if !startTimeVal.IsZero() && !endTimeVal.IsZero() && endTimeVal.Before(startTimeVal) {
		http.Error(w, "ERROR: endTime cannot be before startTime", http.StatusBadRequest)
		return
	}

	dataCars, err := s.carService.GetCarsAvailableInTimeRange(r.Context(), startTimeVal, endTimeVal, string(params.Currency))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	duration := endTimeVal.Sub(startTimeVal).Hours() / 24

	cars := CarIntToListResponse(dataCars, int(duration), params.Currency)

	if err := Util.WriteJSON(w, http.StatusOK, cars); err != nil {
		http.Error(w, "ERROR: failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func (s Server) AddCar(w http.ResponseWriter, r *http.Request) {
	var car data.Car
	if err := Util.DecodeJSON(r, &car); err != nil {
		http.Error(w, "ERROR: invalid request body", http.StatusBadRequest)
		return
	}

	if _, err := s.carService.CreateCar(r.Context(), car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s Server) UpdateCar(w http.ResponseWriter, r *http.Request) {
	var car data.Car
	if err := Util.DecodeJSON(r, &car); err != nil {
		http.Error(w, "ERROR: invalid request body", http.StatusBadRequest)
		return
	}

	_, err := s.carService.UpdateCar(r.Context(), car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	userIdFromRequest, err := getUserIdFromRequest(r)
	if err != nil {
		http.Error(w, "User not found contact support", http.StatusBadRequest)
	}

	user, err := s.userService.GetUserById(r.Context(), userIdFromRequest.String())

	if err := Util.WriteJSON(w, http.StatusOK, MapDataUserToUser(user)); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
	}

}

func (s Server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body UserMutation
	if err := Util.DecodeJSON(r, &body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if body.Username == nil || *body.Username == "" ||
		body.Email == nil || *body.Email == "" ||
		body.Password == nil || *body.Password == "" {
		http.Error(w, "Invalid input: missing fields", http.StatusBadRequest)
		return
	}

	_, err := s.userService.RegisterUser(
		r.Context(),
		data.RentalUser{
			Name:     *body.Username,
			Email:    string(*body.Email),
			Password: *body.Password,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s Server) DeleteUser(w http.ResponseWriter, r *http.Request, params DeleteUserParams) {
	if params.Id == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	err := s.userService.DeleteUser(r.Context(), params.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	dataUsers, err := s.userService.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := make([]User, len(dataUsers))
	for index, db := range dataUsers {
		users[index] = MapDataUserToUser(db)
	}

	if err := Util.WriteJSON(w, http.StatusOK, users); err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
	}
}

func (s Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var body UserMutation
	if err := Util.DecodeJSON(r, &body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if body.Username == nil || *body.Username == "" ||
		body.Email == nil || *body.Email == "" {
		http.Error(w, "Invalid input: missing fields", http.StatusBadRequest)
		return
	}

	userToUpdate, err := s.userService.GetUserByEmail(r.Context(), string(*body.Email))
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	userToUpdate.Name = *body.Username
	if body.Password != nil && *body.Password != "" {
		userToUpdate.Password = *body.Password
	}

	_, err = s.userService.UpdateUser(r.Context(), userToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewServer(dsn string) Server {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	seeder := sql.NewSqlSeeder(db)

	basedOnEnvVaribels(err, seeder)

	use := query.Use(db)

	carRepo := repos.CarRepo{
		Db:  db,
		Ctx: context.Background(),
		Q:   use,
	}

	repo := repos.RentalRepo{
		Q:   use,
		Ctx: context.Background(),
	}

	var userRepo data.UserRepository

	if os.Getenv("DB_BACKEND") == "nosql" {
		nosqldb.InitMongo()
		userRepo = nosqlrepos.NewUserRepo(context.Background(), nosqldb.MongoDatabase)
	} else {
		userRepo = &repos.UserRepo{
			Q:   use,
			Ctx: context.Background(),
		}
	}

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
		slog.Info("Unsupported config no convertor set. Set CONVERTOR_SOAP_URL or CONVERTOR_GRPC_URL if you want an convertor")
	}

	if os.Getenv("PULSAR_LISTENER") == "true" {

		client, err := pulsar.NewClient(pulsar.ClientOptions{
			URL:               os.Getenv("PULSAR_URL"),
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
		if err != nil {
			slog.Error("Could not instantiate Pulsar client: %v", err)
		}

		reader, err := client.CreateReader(pulsar.ReaderOptions{
			Topic:          "car-events",
			StartMessageID: pulsar.EarliestMessageID(),
		})

		if err != nil {
			log.Fatal(err)
		}

		carEvents.NewPulsarConsumer(reader, carRepo)
	}

	return Server{
		carService:     service.NewCarService(carRepo, convService),
		userService:    service.NewUserService(userRepo),
		bookingService: service.NewBookingService(repo, convService),
	}
}

func basedOnEnvVaribels(err error, seeder sql.Seeder) {
	monoSeeder, err := strconv.ParseBool(os.Getenv("SEED_MOMOLITH"))
	if err != nil {
		slog.Info("Not seeding database for Monolith since Argument can't be parsed as boolean")
	}

	if monoSeeder {
		seeder.SeedUser()
		seeder.SeedCars()
		seeder.SeedBooking()
		seeder.SeedBookingMonolithConstraints()
		return
	}

	userSql, err := strconv.ParseBool(os.Getenv("SEED_USER_SQL"))
	if err != nil {
		slog.Info("Not seeding database for SEED_USER_SQL since Argument can't be parsed as boolean")
	}

	if userSql {
		seeder.SeedUser()
		return
	}

	carSQL, err := strconv.ParseBool(os.Getenv("SEED_CAR_SQL"))
	if err != nil {
		slog.Info("Not seeding database for SEED_CAR_SQL since Argument can't be parsed as boolean")
	}

	if carSQL {
		seeder.SeedCars()
	}

	bookSQL, err := strconv.ParseBool(os.Getenv("SEED_BOOKING_SQL"))
	if err != nil {
		slog.Info("Not seeding database for SEED_BOOKING_SQL since Argument can't be parsed as boolean")
	}

	if bookSQL {
		seeder.SeedBooking()
		seeder.SeedBookingConstraints()
	}

}

func MapDataCarToCar(dataCar data.Car) Car {
	price := float32(dataCar.PricePerDay)

	return Car{
		VIN:         &dataCar.Vin,
		Brand:       &dataCar.Brand,
		ImageURL:    &dataCar.ImageUrl,
		Model:       &dataCar.Model,
		PricePerDay: &price,
	}
}

func MapDataUserToUser(user data.RentalUser) User {

	idString := user.Id.String()

	return User{
		Email:    &user.Email,
		Id:       &idString,
		Username: &user.Name,
	}
}

func MapDataBookingToBooking(booking data.Booking) Booking {

	bookingId := booking.Id.String()
	userId := booking.UserId.String()

	f := float32(booking.AmountPaid)

	currency := Currency(booking.Currency)
	return Booking{
		VIN:        &booking.CarVin,
		BookingId:  &bookingId,
		Status:     &booking.Status,
		UserId:     &userId,
		PaidAmount: &f,
		Currency:   &currency,
	}
}

func getUserIdFromRequest(r *http.Request) (uuid.UUID, error) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		return uuid.Max, errors.New("UserId not found in request context")
	}

	return uuid.Parse(userID)

}

func CarIntToListResponse(cars []data.Car, duration int, curr Currency) CarList {

	var carListResponse CarList

	for _, car := range cars {
		pricePerDay := float32(car.PricePerDay)
		priceOverAll := float32(car.PricePerDay * float64(duration))

		carListResponse = append(carListResponse, struct {
			VIN          *string   `json:"VIN,omitempty"`
			Brand        *string   `json:"brand,omitempty"`
			Currency     *Currency `json:"currency,omitempty"`
			ImageURL     *string   `json:"imageURL,omitempty"`
			Model        *string   `json:"model,omitempty"`
			PriceOverAll *float32  `json:"priceOverAll,omitempty"`
			PricePerDay  *float32  `json:"pricePerDay,omitempty"`
		}{
			VIN:          &car.Vin,
			Brand:        &car.Brand,
			Currency:     &curr,
			ImageURL:     &car.ImageUrl,
			Model:        &car.Model,
			PriceOverAll: &priceOverAll,
			PricePerDay:  &pricePerDay,
		})
	}

	return carListResponse
}
