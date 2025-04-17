package data

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetVinsBooked(startTime, endTime time.Time) ([]string, error) {
	endpoint := os.Getenv("BOOKING_SERVICE_URL") + "/bookings/rpc/in_range"

	// Construct the query parameters
	params := url.Values{}

	params.Add("startTime", startTime.Format("2006-01-02"))
	params.Add("endTime", endTime.Format("2006-01-02"))

	fullURL := endpoint
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	// Make the HTTP GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	type Booking struct {
		VIN        *string  `json:"VIN,omitempty"`
		BookingId  *string  `json:"bookingId,omitempty"`
		Currency   *string  `json:"currency,omitempty"`
		PaidAmount *float32 `json:"paidAmount,omitempty"`
		Status     *string  `json:"status,omitempty"`
		UserId     *string  `json:"userId,omitempty"`
	}

	// Decode the JSON response
	var bookingList []Booking
	err = json.NewDecoder(resp.Body).Decode(&bookingList)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	slog.Info("booking list: ", bookingList)

	var result []string
	for _, b := range bookingList {
		result = append(result, *b.VIN)
	}

	return result, nil
}
