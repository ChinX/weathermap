package contorller

import (
	"log"
	"time"

	"github.com/chinx/weathermap/entity"
	"github.com/chinx/weathermap/service/openweather"
	"github.com/chinx/weathermap/utils/cache"
)

func ShowForecastWeather(city string) *entity.ForecastWeatherSummary {
	if city == "" {
		city = "shenzhen,cn"
	}
	return cache.GetOrGenerate(city, getForecastWeather).(*entity.ForecastWeatherSummary)
}

func getForecastWeather(city string) cache.Updater {
	su := &entity.ForecastWeatherSummary{}
	err := openweather.ForecastWeatherSummary(city,
		func(weather *openweather.ForecastWeather) error {
			su.CityName = weather.City.Name
			su.Country = weather.City.Country
			su.CoordinatesLat = weather.City.Coord.Lat
			su.CoordinatesLon = weather.City.Coord.Lon
			su.CurrentTime = int(time.Now().Unix())

			for i := range weather.List {
				item := entity.DateListItem{
					Date:           weather.List[i].DT,
					DateTxt:        weather.List[i].DTTxt,
					Weather:        weather.List[i].Weather[0].Description,
					Image:          weather.List[i].Weather[0].Icon,
					Temperature:    weather.List[i].WMain.Temp,
					TemperatureMax: weather.List[i].WMain.TempMax,
					TemperatureMin: weather.List[i].WMain.TempMin,
					WindSpeed:      weather.List[i].Wind.Speed,
					CloudsDeg:      weather.List[i].Clouds.All,
					Pressure:       weather.List[i].WMain.Pressure,
					Humidity:       weather.List[i].WMain.Humidity,
				}
				if weather.List[i].Rain != nil {
					item.Rain3h = weather.List[i].Rain.Th
				} else {
					item.Rain3h = 0
				}
				su.List = append(su.List, item)
			}
			return nil

		})
	if err != nil {
		log.Println(err)
	}
	return su
}
