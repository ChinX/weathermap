package entity

type CurrentWeatherSummary struct {
	CityName       string  `json:"cityName"`
	Country        string  `json:"country"`
	Temperature    float64 `json:"temperature"`
	Image          string  `json:"image"`
	Date           int     `json:"date"`
	Weather        string  `json:"weather"`
	WindSpeed      float64 `json:"windSpeed"`
	Cloudiness     string  `json:"cloudiness"`
	CloudsDeg      float64 `json:"cloudsDeg"`
	Pressure       float64 `json:"pressure"`
	Humidity       float64 `json:"humidity"`
	Sunrise        int     `json:"sunrise"`
	Sunset         int     `json:"sunset"`
	CoordinatesLat float64 `json:"coordinatesLat"`
	CoordinatesLon float64 `json:"coordinatesLon"`
	UviDate        int     `json:"uviDate"`
	UviDateISO     string  `json:"uviDateISO"`
	UviValue       float64 `json:"uviValue"`
	CurrentTime    int     `json:"currentTime"`
}

func (w *CurrentWeatherSummary) Name() string {
	return w.CityName
}

func (w *CurrentWeatherSummary) UpdateTime() int {
	return w.CurrentTime
}
