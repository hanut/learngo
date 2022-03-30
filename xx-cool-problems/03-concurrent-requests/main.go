package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		panic("Too few arguments. Please pass a comma separated list of cities as the second. if a city name has a space in it use double quotes around the list")
	}
	cs := os.Args[1]
	c := strings.Split(cs, ",")                             // Convert the input to array of city names
	fmt.Println("Getting coordinates of", len(c), "cities") // Print the count of cities
	cm := make(map[string]Coordinates, len(c))              // create a new map with the length of cities
	var wg sync.WaitGroup                                   // new waitgroup to wait for the tasks to finish
	mtx := sync.Mutex{}

	for idx := 0; idx < len(c); idx++ {
		wg.Add(1)
		go func(c string) {
			l, e := GetCoords(c)
			if e != nil {
				panic(errors.Wrap(e, "Error getting location for "+c))
			}
			mtx.Lock()
			cm[strings.ReplaceAll(strings.ToLower(c), " ", "")] = l
			mtx.Unlock()
			wg.Done()
		}(c[idx])
	}

	wg.Wait()

	wd := make(map[string]Weather, len(cm))

	for c, l := range cm {
		wg.Add(1)
		go func(c string, l Coordinates) {
			w, e := GetWeather(l.Lat, l.Lng)
			if e != nil {
				panic(e)
			}
			mtx.Lock()
			wd[c] = w
			mtx.Unlock()
			wg.Done()
		}(c, l)
	}

	wg.Wait()
	sep := "\n" + strings.Repeat("-", 100) + "\n"
	fmt.Println(sep + "Weather data for" + sep)
	mCNSize := getMaxLength(getKeys(wd)) + 4
	for city, data := range wd {
		fmt.Printf("%s(⇡: %07.2f) is %05.2f °C with winds of speed %04.2f & direction %04.2f\n", padString(strings.ToUpper(city), mCNSize), data.Elevation, data.Temperature, data.WindSpeed, data.WindDirection)
	}
}
func getKeys(m map[string]Weather) []string {
	keys := make([]string, len(m))
	var counter int
	for k := range m {
		keys[counter] = k
		counter++
	}
	return keys
}

func getMaxLength(s []string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		if len(s[i]) > max {
			max = len(s[i])
		}
	}
	return max
}

func padString(s string, size int) string {
	l := len(s)
	if l == size {
		return s
	}

	return s + strings.Repeat(" ", size-l)
}
