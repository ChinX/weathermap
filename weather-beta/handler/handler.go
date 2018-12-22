package handler

import (
	"encoding/json"
	"github.com/chinx/weathermap/weather-beta/contorller"
	"net/http"
)

func ShowWeather(w http.ResponseWriter, r *http.Request)  {
	city := r.URL.Query().Get("city")
	data := contorller.ShowCurrentWeather(city)
	byteData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(byteData)
}
