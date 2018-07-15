package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const appID = ""

func main() {
	city := os.Args
	var cityName string
	if len(city) > 1 {
		cityName = city[1]
	} else {
		fmt.Println("Please provide a city name as Command line arg")
		os.Exit(0)
	}

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
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + cityName + "&units=metric&appid=" + appID)
	if err != nil {
		fmt.Println("Error while fetching weather information")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(0)
	}
	w := weatherInfo{}
	json.Unmarshal(contents, &w)
	if w.Name != "" {
		fmt.Printf("Name -> %s \nTemperature -> %d C \nDescription -> %s\nHumidity %d\n", w.Name, w.Main.Temp, w.Weather[0].Description, w.Main.Humidity)
	} else {
		fmt.Printf("No weather information found for the city '%s' \n", cityName)
	}
}
