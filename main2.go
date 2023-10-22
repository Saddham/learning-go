package main

import (
	"fmt"
	"learning-go/models"
	"learning-go/printer"
)

func main2() {
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	lon := models.NewCity("London", 7., true, false)

	ams := models.NewCity("Amsterdam", 11, false, true)

	nyc := models.NewCity("New York", -3, true, true)

	p.CityDetails(lon)
	p.CityDetails(ams)
	p.CityDetails(nyc)

	// Arrays
	cities := [2]string{"London"}
	copy := cities

	cities[1] = "New York"
	addCityPtr("Miami", &copy)

	fmt.Println("Cities: ", cities)
	fmt.Println("Cities Copy: ", copy)

	// Slices (Dynamic Arrays)
	cities2 := make([]string, 1, 1)
	cities2[0] = "Amsterdam"            // Set value at index 0
	cities2 = append(cities2, "London") // Append to to the end of slice (after initialized length), irrespective of empty values at the start of the slice
	cities2 = append(cities2, "New York")

	fmt.Println("Cities2 Slice1: ", cities2[0:1])
	fmt.Println("Cities2 Slice:2 ", cities2[1:])
	fmt.Println("Cities2 Slice3: ", cities2[:2])

	// Maps
	map1 := make(map[string]int)
	map1["London"] = 10

	fmt.Printf("Map[%v]= %v\n", "London", map1["London"])

	delete(map1, "London")

	_, present := map1["London"]
	fmt.Printf("Map[%v]= %v\n", "London", present)

	// Iterating
	for i := 0; i < len(cities); i++ { // For loop
		fmt.Printf("[%v]=%v\n", i, cities[i])
	}

	for i, value := range cities { // Range operator
		fmt.Printf("[%v]=%v\n", i, value)
	}
}

func addCity(city string, cities [2]string) {
	cities[1] = city
}

func addCityPtr(city string, cities *[2]string) {
	cities[1] = city
}
