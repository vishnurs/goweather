package main

import (
	"fmt"
	"os"
	"sync"
)

const appID = "fbb35ad5307e2ce97c0d2763399b83a3"

type weatherInfo struct {
	Weather []struct {
		Main        string
		Description string
	}
	Main struct {
		Temp     int8
		Humidity int8
	}
	Name string
}

func main() {

	var wg sync.WaitGroup
	args := os.Args
	cities := args[1:]
	if len(cities) == 0 {
		fmt.Println("Please provide a city name as Command line arg")
		os.Exit(0)
	}
	for _, city := range cities {
		wg.Add(1)
		go getWeather(city, &wg)
	}
	wg.Wait()

}
