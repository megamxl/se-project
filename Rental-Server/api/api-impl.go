package api

import (
	"encoding/json"
	"github.com/megamxl/se-project/Rental-Server/api/DTO"
	"github.com/megamxl/se-project/Rental-Server/api/Util"
	"github.com/megamxl/se-project/Rental-Server/internal/service"
	"net/http"
	"time"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct {
	carService     service.CarService
	userService    service.UserService
	bookingService service.BookingService
}

func (s Server) Login(w http.ResponseWriter, r *http.Request) {
	resp := TokenResponse{Token: ptr("eyMyToken")}
	w.Header().Set("Content-Type", "application/json") // Set header first
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
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
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "missing userId query parameter", http.StatusBadRequest)
		return
	}

	bookings, err := s.bookingService.GetAllBookingsByUser(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := Util.WriteJSON(w, http.StatusOK, bookings); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func (s Server) BookCar(w http.ResponseWriter, r *http.Request) {
	var req DTO.Booking
	if err := Util.DecodeJSON(r, &req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	_, err := s.bookingService.BookCar(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	if err := Util.WriteJSON(w, http.StatusOK, booking); err != nil {
		http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func (s Server) GetAllBookingsByUser(w http.ResponseWriter, r *http.Request) {
	bookings, err := s.bookingService.GetAllBookings(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
		startTimeVal = *params.StartTime
	}

	if params.EndTime != nil {
		endTimeVal = *params.EndTime
	}

	if !startTimeVal.IsZero() && !endTimeVal.IsZero() && endTimeVal.Before(startTimeVal) {
		http.Error(w, "ERROR: endTime cannot be before startTime", http.StatusBadRequest)
		return
	}

	cars, err := s.carService.GetCarsAvailableInTimeRange(r.Context(), startTimeVal, endTimeVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := Util.WriteJSON(w, http.StatusOK, cars); err != nil {
		http.Error(w, "ERROR: failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func (s Server) AddCar(w http.ResponseWriter, r *http.Request) {
	var car DTO.Car
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
	var car DTO.Car
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
	//TODO Implementation in future PR
	panic("implement me")
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
		DTO.RentalUser{
			Name:  *body.Username,
			Email: string(*body.Email),
		},
		*body.Password,
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
	users, err := s.userService.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

func NewServer() Server {
	return Server{}
}

func ptr(s string) *string {
	return &s
}
