package main

import (
	"github.com/chinx/weathermap/forecast/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/forecast/show", handler.ShowWeather)
	http.ListenAndServe(":8081", nil)
}
