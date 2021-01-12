package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	_, err = csvReader.Read()
	if err != nil {
		log.Fatal("Unable to read csv header "+filePath, err)
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

// ParseCsv : Parse csv events
func ParseCsv(filePath string) {
	records := readCsvFile("./events.csv")
	for index, record := range records {
		CreateUnit(index, record)
	}
}
