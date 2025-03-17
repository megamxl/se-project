package api

import (
	"encoding/json"
	"net/http"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{}

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
	//TODO implement me
	panic("implement me")
}

func (s Server) ListCars(w http.ResponseWriter, r *http.Request, params ListCarsParams) {
	//TODO implement me
	panic("implement me")
}

func (s Server) AddCar(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s Server) UpdateCar(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
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
