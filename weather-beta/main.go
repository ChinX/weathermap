package main

import (
	"github.com/chinx/weathermap/weather-beta/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/weather/show", handler.ShowWeather)
	http.ListenAndServe(":8080", nil)
}
