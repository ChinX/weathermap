package main

import (
	"net/http"

	"github.com/chinx/weathermap/fusionweather/handler"
)

func main() {
	http.HandleFunc("/fusionweather/show", handler.ShowWeather)
	http.ListenAndServe(":8082", nil)
}
