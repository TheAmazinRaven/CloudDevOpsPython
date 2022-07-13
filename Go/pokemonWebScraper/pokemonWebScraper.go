package main

import (
	/* "encoding/csv"
	"encoding/json" */
	"fmt"

	"github.com/gocolly/colly"
	/* "log"
	"os"
	"strconv" */)

type pokemon struct {
	Name   string `json:"Name"`
	Type   string `json:"Type"`
	ImgUrl string `json:"imgurl"`
}

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("https://bulbapedia.bulbagarden.net"),
	)

	c.OnHTML("", func(h *colly.HTMLElement) { // i have no idea where to pull the data from?
		fmt.Println("Pokemon Name " + h.Text) // need to change this to the header element in the HTML

	})

	c.Visit("https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_National_Pok%C3%A9dex_number")
}
