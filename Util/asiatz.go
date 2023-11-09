package Util

import (
	"fmt"
	"strconv"
)

// ToUTC converts a time string (HH:mm) from a specified timezone to UTC time string (HH:mm).
func ToUTC(timezoneOffset float64, time string) (string, error) {
	hour, err := strconv.Atoi(time[:2])
	if err != nil {
		return "", err
	}
	minute, err := strconv.Atoi(time[3:])
	if err != nil {
		return "", err
	}
	totalMinutes := hour*60 + minute
	utcTotalMinutes := ((totalMinutes-int(timezoneOffset*60))%1440 + 1440) % 1440
	utcHour := utcTotalMinutes / 60
	utcMinute := utcTotalMinutes % 60
	utcTime := fmt.Sprintf("%02d:%02d", utcHour, utcMinute)
	return utcTime, nil
}

// ShanghaiToUTC converts a Shanghai time string (HH:mm) to UTC time string (HH:mm).
// For example, "08:00" in Shanghai is equivalent to "00:00" in UTC.
func ShanghaiToUTC(shanghaiTime string) (string, error) {
	return ToUTC(8, shanghaiTime)
}
