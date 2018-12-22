package openweather

type CurrentWeather struct {
	Coord   Coord       `json:"coord"`
	Weather []*Weather  `json:"weather"`
	Base    string      `json:"base"`
	WMain   WeatherMain `json:"main"`
	Wind    Wind        `json:"wind"`
	Rain    CurrentRain `json:"rain"`
	Clouds  Clouds      `json:"clouds"`
	DT      int         `json:"dt"`
	WSys    CurrentSys  `json:"sys"`
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Cod     int         `json:"cod"`
}

type ForecastWeather struct {
	City    City       `json:"city"`
	Cnt     int        `json:"cnt"`
	Cod     string     `json:"cod"`
	Message float64    `json:"message"`
	List    []ListItem `json:"list"`
}

type ListItem struct {
	DT      int           `json:"dt"`
	DTTxt   string        `json:"dt_txt"`
	Weather []Weather     `json:"weather"`
	WMain   WeatherMain   `json:"main"`
	Clouds  Clouds        `json:"clouds"`
	WSys    ForecastSys   `json:"sys"`
	Wind    Wind          `json:"wind"`
	Rain    *ForecastRain `json:"rain"`
}

type City struct {
	Country string `json:"country"`
	Coord   Coord  `json:"coord"`
	Name    string `json:"name"`
	ID      int    `json:"id"`
}

type Clouds struct {
	All float64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	MainStr     string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherMain struct {
	Temp      float64 `json:"temp"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	GrndLevel float64 `json:"grnd_level"`
	TempKf    float64 `json:"temp_kf"`
	SeaLevel  float64 `json:"sea_level"`
}

type CurrentRain struct {
	Oh float64 `json:"1h"`
}

type ForecastRain struct {
	Th float64 `json:"3h"`
}

type CurrentSys struct {
	Kind    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

type ForecastSys struct {
	Pod string `json:"pod"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type UltravioletIndex struct {
	Date    int     `json:"date"`
	DateISO string  `json:"date_iso"`
	Lon     float64 `json:"lon"`
	Lat     float64 `json:"lat"`
	Value   float64 `json:"value"`
}
