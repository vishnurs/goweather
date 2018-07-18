package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func getWeather(city string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&appid=" + appID)
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
		fmt.Printf("Name -> %s \nTemperature -> %d C \nDescription -> %s\nHumidity %d\n\n", w.Name, w.Main.Temp, w.Weather[0].Description, w.Main.Humidity)
		fmt.Println("------------------------------------------")
	} else {
		fmt.Printf("No weather information found for the city '%s' \n", city)
	}
}
