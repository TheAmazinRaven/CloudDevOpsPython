// https://webdamn.com/write-data-to-csv-file-using-golang/

package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	empData := [][]string{
		{"Name", "City", "Skills"},
		{"Smith", "New York", "Java"},
		{"William", "Paris", "Golang"},
		{"Rose", "London", "PHP"},
	}

	csvFile, err := os.Create("employee.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()
	csvFile.Close()

}
