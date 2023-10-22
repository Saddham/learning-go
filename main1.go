package main

import (
	"fmt"
)

var (
	v0 int32
	v1 float64
	v2 bool
	v3 string
)

func main1() {
	// var greeting = "Hello, World!"
	// Greeting := "Hello, World!" => public variable
	greeting := "Hello, World!" // Initialization with declaration

	fmt.Println(greeting)

	fmt.Printf("V0: Type - %T, Value - %v\n", v0, v0)
	fmt.Printf("V1: Type - %T, Value - %v\n", v1, v1)
	fmt.Printf("V2: Type - %T, Value - %v\n", v2, v2)
	fmt.Printf("V3: Type - %T, Value - %v\n", v3, v3)

	v3 = "Pointer value"

	var ptr *string

	ptr = &v3

	fmt.Println("Address: ", ptr)
	fmt.Println("Value: ", *ptr)

	// Multiple return values
	area, circumf := rectangle(2, 3)
	fmt.Printf("rectangle(%v, %v): area = %v;\n", 2, 3, area)
	fmt.Printf("rectangle(%v, %v): circumf = %v;\n", 2, 3, circumf)

	// Multiple return values with named variables
	area, _ = named_rectangle(5, 10)
	fmt.Printf("named_rectangle(%v, %v): area = %v;\n", 5, 10, area)

	// If else
	printParity(2)
	printParity(3)

	// Error handling
	x := 5
	y := 10
	area_err, err := area_with_error(x, y)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("area_with_error(%v, %v): area = %v;\n", x, y, *area_err)

	// Deferred functions
	msg := "Hello, world!"
	defer printMsg("Defer-0", msg)

	msg = "Here, world!"
	defer printMsg("Defer-1", msg)

	msg = "Goodbye, world!"
	fmt.Println("Main has exited")

	// Structs
	c := city{
		name: "London",
	}

	fmt.Println(c.name)
	fmt.Println(newCity("Pune").name)

	// Adding methods to structs
	c.tempC = 7.5
	tempF := c.tempF()

	fmt.Printf("[%v]: tempC:%v; tempF:%v\n", c.name, c.tempC, tempF)
}

func rectangle(x int, y int) (int, int) {
	area := x * y
	circumf := 2 * (x + y)

	return area, circumf
}

func named_rectangle(x int, y int) (area int, circumf int) {
	area = x * y
	circumf = 2 * (x + y)

	return
}

func printParity(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}

	fmt.Printf("%v is odd.\n", x)
}

func area_with_error(x int, y int) (*int, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero area: [%v, %v]", x, y)
	}

	area := x * y

	return &area, nil
}

func printMsg(id string, m string) {
	fmt.Printf("[%v]: %v\n", id, m)
}

type city struct {
	name  string
	tempC float64
}

func newCity(n string) city {
	return city{
		name: n,
	}
}

func (c city) tempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func (c *city) tempCF() {
	c.tempC = (c.tempC * 9 / 5) + 32
}
