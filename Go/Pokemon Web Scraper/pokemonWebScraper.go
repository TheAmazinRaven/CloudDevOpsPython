package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func main() {

	// creating file
	fName := "data.csv"
	// logging errors
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err :%q", err)
		return
	}

	// closing the file

	defer file.Close()

	// writing to the CSV file

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.newCollector(
		colly.AllowedDomains("https://bulbapedia.bulbagarden.net"),
	)

}
