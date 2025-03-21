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
	carService service.CarService
}

func (s Server) Login(w http.ResponseWriter, r *http.Request) {
	resp := TokenResponse{Token: ptr("eyMyToken")}
	w.Header().Set("Content-Type", "application/json") // Set header first
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) DeleteBooking(w http.ResponseWriter, r *http.Request, params DeleteBookingParams) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetBookings(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s Server) BookCar(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s Server) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetBookingById(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetAllBookingsByUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (s Server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s Server) DeleteUser(w http.ResponseWriter, r *http.Request, params DeleteUserParams) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewServer() Server {
	return Server{}
}

func ptr(s string) *string {
	return &s
}
