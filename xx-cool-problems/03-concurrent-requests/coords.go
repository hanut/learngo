package main

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type Coordinates struct {
	Lat string
	Lng string
}

type City struct {
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}

var cities = make(map[string]Coordinates)

func init() {
	// read the cities json file
	d, e := os.ReadFile("cities.json")
	if e != nil {
		panic(e)
	}
	var tcl []City
	e = json.Unmarshal(d, &tcl)
	if e != nil {
		panic(e)
	}
	for _, c := range tcl {
		cities[c.Name] = Coordinates{Lat: c.Lat, Lng: c.Lng}
	}
}

func GetCoords(city string) (Coordinates, error) {
	var loc Coordinates
	loc, exists := cities[city]
	if !exists {
		return loc, errors.New("unsupported city")
	}
	return loc, nil
}
