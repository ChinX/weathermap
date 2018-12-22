package openweather

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chinx/weathermap/utils/rest"
)

const (
	appKey      = "763d8bb819e1b0fb58c8385ddd26856e"
	defaultCity = "ShenZhen,CN"
	weather     = "weather"
	forecast    = "forecast"
)

var (
	URL_HTTPS = "https://api.openweathermap.org/data/2.5/%s?appid=%s&q=%s&units=metric"
	URL_HTTP  = "http://api.openweathermap.org/data/2.5/%s?appid=%s&q=%s&units=metric"

	UVI_URL_HTTP  = "http://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%F&lon=%F"
	UVI_URL_HTTPS = "https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%F&lon=%F"
)

func CurrentWeatherSummary(city string, transform func(*CurrentWeather) error) error {
	data := &CurrentWeather{}
	err := restGet(fmt.Sprintf(URL_HTTP, weather, appKey, formatCity(city)), data)
	if err == nil {
		err = transform(data)
	}
	return err
}

func CurrentWeatherUltravioletIndex(lat, lon float64, transform func(*UltravioletIndex) error) error {
	data := &UltravioletIndex{}
	err := restGet(fmt.Sprintf(fmt.Sprintf(UVI_URL_HTTP, appKey, lat, lon)), data)
	if err == nil {
		err = transform(data)
	}
	return err
}

func ForecastWeatherSummary(city string, transform func(*ForecastWeather) error) error {
	data := &ForecastWeather{}
	err := restGet(fmt.Sprintf(URL_HTTP, forecast, appKey, formatCity(city)), data)
	if err == nil {
		err = transform(data)
	}
	return err
}

func restGet(rawURL string, v interface{}) error {
	resp, err := rest.Get(rawURL, nil, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	if err = rest.ParseResponse(resp, http.StatusOK, v); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func formatCity(city string) string {
	if city == "" {
		return defaultCity
	}
	return city
}
