package main

import (
	"fmt"
	// "os"
)

const (
	// UnitPrefix : Unit files prefix for this app
	UnitPrefix = "sn" // systemd notifier
)

func main() {
	// if len(os.Args[1]) != 2 {
	// 	Usage(os.Args[0])
	// 	os.Exit((1))
	// }
	DeleteAllUnits()
	// ParseCsv(os.Args[1])
	ParseCsv("./events.csv")
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "All events imported successfully!")
}
