package contorller

import (
	"log"
	"time"

	"github.com/chinx/weathermap/entity"
	"github.com/chinx/weathermap/service/openweather"
	"github.com/chinx/weathermap/utils/cache"
)

func ShowCurrentWeather(city string) *entity.CurrentWeatherSummary {
	if city == "" {
		city = "shenzhen,cn"
	}
	return cache.GetOrGenerate(city, getCurrentWeather).(*entity.CurrentWeatherSummary)
}

func getCurrentWeather(city string) cache.Updater {
	su := &entity.CurrentWeatherSummary{}
	err := openweather.CurrentWeatherSummary(city, func(weather *openweather.CurrentWeather) error {
		su.CityName = weather.Name
		su.Country = weather.WSys.Country
		su.Temperature = weather.WMain.Temp
		su.Image = weather.Weather[0].Icon
		su.Date = weather.DT
		su.Weather = weather.Weather[0].Description
		su.WindSpeed = weather.Wind.Speed
		su.Cloudiness = weather.Weather[0].Description
		su.CloudsDeg = weather.Clouds.All
		su.Pressure = weather.WMain.Pressure
		su.Humidity = weather.WMain.Humidity
		su.Sunrise = weather.WSys.Sunrise
		su.Sunset = weather.WSys.Sunset
		su.CoordinatesLat = weather.Coord.Lat
		su.CoordinatesLon = weather.Coord.Lon
		su.CurrentTime = int(time.Now().Unix())
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return su
}
