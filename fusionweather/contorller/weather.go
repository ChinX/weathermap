package contorller

import (
	"log"
	"net/http"

	"github.com/chinx/weathermap/entity"
	"github.com/chinx/weathermap/utils/rest"
)

func ShowFusionWeather(city string) *entity.FusionWeatherSummary {
	return &entity.FusionWeatherSummary{
		CurrentWeather:  achieveCurrentWeatherSummary(city),
		ForecastWeather: achieveForecastWeatherSummary(city),
	}
}

func achieveCurrentWeatherSummary(city string) entity.CurrentWeatherSummary {
	currentURL := "http://weather:8080/weather/show?city=" + city
	current := entity.CurrentWeatherSummary{}

	resp, err := rest.Get(currentURL, nil, nil)
	if err == nil {
		err = rest.ParseResponse(resp, http.StatusOK, &current)
	}

	if err != nil {
		log.Println(err)
	}
	return current
}

func achieveForecastWeatherSummary(city string) entity.ForecastWeatherSummary {
	currentURL := "http://forecast:8081/forecast/show?city=" + city
	forecast := entity.ForecastWeatherSummary{}

	resp, err := rest.Get(currentURL, nil, nil)
	if err == nil {
		err = rest.ParseResponse(resp, http.StatusOK, &forecast)
	}

	if err != nil {
		log.Println(err)
	}
	return forecast
}
