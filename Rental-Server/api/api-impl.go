package api

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/megamxl/se-project/Rental-Server/api/Util"
	"github.com/megamxl/se-project/Rental-Server/internal/di"
	"strings"

	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/middleware"
	"github.com/megamxl/se-project/Rental-Server/internal/service"
	"log"
	"log/slog"
	"net/http"
	"os"
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

func (s Server) GetCarByVin(w http.ResponseWriter, r *http.Request, params GetCarByVinParams) {

	vin, err := s.carService.GetCarByVin(context.Background(), params.VIN)

	if err != nil {
		http.Error(w, "No Car With that vin", http.StatusUnauthorized)
		return
	}

	car := Car{
		VIN:         &vin.Vin,
		Brand:       &vin.Brand,
		ImageURL:    &vin.ImageUrl,
		Model:       &vin.Model,
		PricePerDay: nil,
	}

	if err := Util.WriteJSON(w, http.StatusOK, car); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusBadRequest)
		return
	}

}

func (s Server) ListBookingsInRange(w http.ResponseWriter, r *http.Request, params ListBookingsInRangeParams) {
	timeRange, err := s.bookingService.GetAllBookingsInTimeRange(params.StartTime.Time, params.EndTime.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookings := make([]Booking, len(timeRange))
	for index, db := range timeRange {
		bookings[index] = MapDataBookingToBooking(db)
	}

	if err := Util.WriteJSON(w, http.StatusOK, bookings); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusBadRequest)
		return
	}
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

// Bookings

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
		return
	}

	vin, err := s.carService.GetCarByVin(r.Context(), *req.VIN)
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
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
		http.Error(w, "Failed to create booking: "+err.Error(), http.StatusUnprocessableEntity)
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

	if err := data.ValidateRequiredFields(req, []string{"BookingId", "Status"}); err != nil {
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

	if err := data.ValidateRequiredFields(struct{ ID string }{ID: id}, []string{"ID"}); err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	booking, err := s.bookingService.GetBookingById(r.Context(), id)
	if err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	if err := Util.WriteJSON(w, http.StatusOK, MapDataBookingToBooking(booking)); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func (s Server) GetAllBookingsByUser(w http.ResponseWriter, r *http.Request) {
	userIdFromRequest, err := getUserIdFromRequest(r)
	if err != nil {
		http.Error(w, "User not found in context", http.StatusBadRequest)
		return
	}
	dataBookings, err := s.bookingService.GetAllBookingsByUser(r.Context(), userIdFromRequest.String())
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

// Cars

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

	startDate := time.Date(startTimeVal.Year(), startTimeVal.Month(), startTimeVal.Day(), 0, 0, 0, 0, startTimeVal.Location())
	endDate := time.Date(endTimeVal.Year(), endTimeVal.Month(), endTimeVal.Day(), 0, 0, 0, 0, endTimeVal.Location())

	duration := int(endDate.Sub(startDate).Hours()/24) + 1

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

	if err := data.ValidateRequiredFields(car, []string{"Vin", "Model", "Brand", "PricePerDay"}); err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusUnprocessableEntity)
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

// Users

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

	request, err2 := getUserIdFromRequest(r)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
	}

	userToUpdate, err := s.userService.GetUserById(r.Context(), request.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	allEmpty := true

	if body.Username != nil && *body.Username != "" {
		userToUpdate.Name = *body.Username
		allEmpty = false
	}

	if body.Password != nil && *body.Password != "" {
		userToUpdate.Password = *body.Password
		allEmpty = false
	}

	if body.Email != nil && *body.Email != "" {
		userToUpdate.Email = string(*body.Email)
		allEmpty = false
	}

	if allEmpty {
		http.Error(w, "Invalid input: no fields would change", http.StatusBadRequest)
	}

	_, err = s.userService.UpdateUser(r.Context(), userToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewServer() Server {

	dbBackend := os.Getenv("DB_BACKEND")

	var userRepo data.UserRepository
	var carRepo data.CarRepository
	var bookingRepo data.BookingRepository

	switch dbBackend {
	case "SQL":
		slog.Info("DB_BACKEND=SQL")
		connection := di.GetSQLDatabaseConnection()
		di.SeedDatabase(connection)

		userRepo = di.GetUserRepositorySQL(connection)
		carRepo = di.GetCarRepositorySQL(connection)
		bookingRepo = di.GetBookingRepositorySQL(connection)

	case "NO-SQL":
		nosqldb.InitMongoWith(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB_NAME"))
		carRepo = nosqlrepos.NewCarRepo(context.Background(), nosqldb.MongoDatabase)
		userRepo = nosqlrepos.NewUserRepo(context.Background(), nosqldb.MongoDatabase)

		slog.Info("DB_BACKEND=NO-SQL")
	default:
		log.Fatal("DB_BACKEND= Not Set no function reconfigure the app")
	}

	if userRepo == nil {
		slog.Info("🚨 UserRepo is nil")
	}

	if carRepo == nil {
		slog.Info("🚨 CarRepo is nil")
	}

	if bookingRepo == nil {
		slog.Info("🚨 BookingRepo is nil")
	}

	convertor := di.GetConvertor()

	var err error

	for i := 0; i < 5; i++ {
		err = nil
		err = di.PulsarListner(carRepo)
		if err == nil {
			break
		}
		slog.Info("Sleeping 5 seconds and then trying to connect to Pulsar Again")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("🔌 Cant connect to pulsar", err)
	}

	if os.Getenv("ADMIN") != "" {
		split := strings.Split(os.Getenv("ADMIN"), ";")

		userRepo.SaveUser(
			data.RentalUser{
				Id:       uuid.UUID{},
				Name:     "admin",
				Email:    split[0],
				Password: split[1],
				Admin:    true,
			})

	}

	return Server{
		carService:     service.NewCarService(carRepo, convertor),
		userService:    service.NewUserService(userRepo),
		bookingService: service.NewBookingService(bookingRepo, convertor),
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

	s := booking.StartTime.Format("2006-01-02")
	e := booking.EndTime.Format("2006-01-02")

	currency := Currency(booking.Currency)
	return Booking{
		VIN:        &booking.CarVin,
		BookingId:  &bookingId,
		Status:     &booking.Status,
		UserId:     &userId,
		PaidAmount: &f,
		Currency:   &currency,
		StartDate:  &s,
		EndDate:    &e,
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
