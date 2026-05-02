package routes 

import (
	"net/http"
)
func V1Routes(mux *http.ServeMux){
	mux.HandleFunc("GET /health",health);
	mux.HandleFunc("GET /weatherData",getWeather)
	mux.HandleFunc("GET /airQualityData",getAirQuality)
}