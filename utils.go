package main

import (
	"log"
	"strconv"
)

// ConvertCalendarToInt : Cast year, month, day from string to int
func ConvertCalendarToInt(year string, month string, day string) (int, int, int) {
	y, err := strconv.Atoi(year)
	if err != nil {
		log.Fatal("Cannot cast year to integer", err)
	}
	m, err := strconv.Atoi(month)
	if err != nil {
		log.Fatal("Cannot cast month to integer", err)
	}
	d, err := strconv.Atoi(day)
	if err != nil {
		log.Fatal("Cannot cast day to integer", err)
	}
	return y, m, d
}

// ConvertCalendarToString : Cast year, month, day from int to string
func ConvertCalendarToString(year int, month int, day int) (string, string, string) {
	y := strconv.Itoa(year)
	m := strconv.Itoa(month)
	d := strconv.Itoa(day)
	return y, m, d
}
