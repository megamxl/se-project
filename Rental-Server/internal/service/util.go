package service

import "time"

func GetDurationBetween(startTimeVal time.Time, endTimeVal time.Time) int {
	startDate := time.Date(startTimeVal.Year(), startTimeVal.Month(), startTimeVal.Day(), 0, 0, 0, 0, startTimeVal.Location())
	endDate := time.Date(endTimeVal.Year(), endTimeVal.Month(), endTimeVal.Day(), 0, 0, 0, 0, endTimeVal.Location())

	return int(endDate.Sub(startDate).Hours()/24) + 1
}
