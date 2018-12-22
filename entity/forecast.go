package entity

type ForecastWeatherSummary struct {
	CityName       string         `json:"cityName"`
	Country        string         `json:"country"`
	CoordinatesLat float64        `json:"coordinatesLat"`
	CoordinatesLon float64        `json:"coordinatesLon"`
	List           []DateListItem `json:"dateList"`
	CurrentTime    int            `json:"currentTime"`
}

type DateListItem struct {
	Date           int     `json:"date"`
	Image          string  `json:"image"`
	DateTxt        string  `json:"dateTxt"`
	TemperatureMax float64 `json:"temperatureMax"`
	Temperature    float64 `json:"temperature"`
	TemperatureMin float64 `json:"temperatureMin"`
	Weather        string  `json:"weather"`
	Pressure       float64 `json:"pressure"`
	Humidity       float64 `json:"humidity"`
	WindSpeed      float64 `json:"windSpeed"`
	CloudsDeg      float64 `json:"cloudsDeg"`
	Rain3h         float64 `json:"rain3h"`
}

func (w *ForecastWeatherSummary) Name() string {
	return w.CityName
}

func (w *ForecastWeatherSummary) UpdateTime() int {
	return w.CurrentTime
}
