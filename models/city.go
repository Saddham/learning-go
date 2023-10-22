package models

type city struct {
	name     string
	tempC    float64
	beach    bool
	mountain bool
}

type ICity interface {
	Name() string
	TempC() float64
	TempF() float64
	HasBeach() bool
	HasMountain() bool
}

func NewCity(n string, t float64, b bool, m bool) ICity {
	return &city{
		name:     n,
		tempC:    t,
		beach:    b,
		mountain: m,
	}
}

func (c city) Name() string {
	return c.name
}

func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func (c city) TempC() float64 {
	return c.tempC
}

func (c city) HasBeach() bool {
	return c.beach
}

func (c city) HasMountain() bool {
	return c.mountain
}
