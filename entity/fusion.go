package entity

type FusionWeatherSummary struct {
	CurrentWeather  CurrentWeatherSummary  `json:"currentWeather"`
	ForecastWeather ForecastWeatherSummary `json:"forecastWeather"`
}
