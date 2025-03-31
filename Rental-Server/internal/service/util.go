package service

import "time"

func GetDurationBetween(start time.Time, end time.Time) time.Duration {
	after := end.After(start)
	if !after {
		return time.Duration(0)
	}

	return end.Sub(start)
}
