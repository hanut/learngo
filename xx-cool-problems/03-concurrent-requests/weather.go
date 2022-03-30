package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const ApiUrl = "https://api.open-meteo.com/v1/forecast?"

type Weather struct {
	WindSpeed     float32
	WindDirection float32
	Elevation     float32
	Temperature   float32
}

func GetWeather(lat string, lng string) (Weather, error) {
	var w Weather

	uri := fmt.Sprintf("%slatitude=%s&longitude=%s&current_weather=true", ApiUrl, lat, lng)

	res, err := http.Get(uri)

	if err != nil {
		return w, errors.Wrap(err, "Failed to retrieve weather for "+lat+","+lng)
	}

	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return w, errors.Wrap(err, "Failed to read response for "+lat+","+lng)
	}
	var t struct {
		Elevation float32 `json:"elevation,omitempty"`
		Weather   struct {
			Temperature   float32 `json:"temperature,omitempty"`
			WindSpeed     float32 `json:"windspeed,omitempty"`
			WindDirection float32 `json:"winddirection,omitempty"`
		} `json:"current_weather,omitempty"`
	}
	err = json.Unmarshal(d, &t)
	if err != nil {
		return w, errors.Wrap(err, "Failed to parse json response for "+lat+","+lng)
	}
	w.Elevation = t.Elevation
	w.Temperature = t.Weather.Temperature
	w.WindDirection = t.Weather.WindDirection
	w.WindSpeed = t.Weather.WindSpeed

	return w, nil
}
