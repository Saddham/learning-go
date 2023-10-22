package models

import (
	"learning-go/data"
	"sort"
)

type cities struct {
	cityMap map[string]ICity
}

type ICities interface {
	ListAll() []ICity
	Filter(beach bool, ski bool) []ICity
}

func NewCities(reader data.DataReader) (ICities, error) {
	d, err := reader.ReadData()

	if err != nil {
		return nil, err
	}

	cityMap := make(map[string]ICity)

	for _, r := range d {
		cityMap[r.Id] = NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountain)
	}

	return &cities{
		cityMap,
	}, nil
}

func (c cities) ListAll() []ICity {
	var cs []ICity

	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}

	sortAlphabetically(cs)

	return cs
}

func sortAlphabetically(cs []ICity) {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Name() < cs[j].Name()
	})
}

func (c cities) Filter(beach bool, ski bool) []ICity {
	if !beach && !ski {
		return c.ListAll()
	}

	var cs []ICity

	for _, rc := range c.cityMap {
		if matchFilter(rc, beach, ski) {
			cs = append(cs, rc)
		}
	}

	sortAlphabetically(cs)

	return cs
}

func matchFilter(rc ICity, beach bool, ski bool) bool {
	if beach && rc.HasBeach() {
		return true
	}

	if ski && rc.HasMountain() {
		return true
	}

	return false
}
