package main

import (
	"flag"
	"fmt"
	"learning-go/data"
	"learning-go/models"
	"learning-go/printer"
)

func main() {
	fmt.Printf("Welcome to the Temperature Service!")

	beachReady := flag.Bool("beach", false, "Display only beach ready destinations")
	skiReady := flag.Bool("ski", false, "Display only ski ready destinations")
	flag.Parse()

	cities, err := models.NewCities(data.NewReader())

	if err != nil {
		fmt.Println("Fatal error occurred: ", err)
		return
	}

	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	cs := cities.Filter(*beachReady, *skiReady)

	for _, c := range cs {
		p.CityDetails(c)
	}

	// Generics
	dest := []string{"London", "Paris", "New York", "Milan", "Sydney"}
	locs := []location{
		{name: "Tokya", temp: 34},
		{name: "Concun", temp: 9},
	}

	print(dest)
	print(locs)

	// Concurrency
	dc := make(chan string)
	go writeDestinations(dc)

	for d := range dc {
		fmt.Printf("[Main]: Hello from %v!\n", d)
	}

	fmt.Println("[Main]: Shutting down.")
}

func writeDestinations(dc chan string) {
	dest := []string{"London", "Paris", "New York", "Milan", "Sydney"}

	for _, d := range dest {
		fmt.Printf("[Writer]: Writing %v to channel.\n", d)

		dc <- d
	}

	close(dc)

	fmt.Println("[Writer]: Shutting down.")
}

type location struct {
	name string
	temp int
}

func print[T any](in []T) {
	for _, t := range in {
		fmt.Println(t)
	}
}
