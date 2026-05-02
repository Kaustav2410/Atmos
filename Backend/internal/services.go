package routes
// TODO: Create a helper function so that services code won't be written again and again
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
var logger = log.Default()
type weatherData struct{
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone string `json:"timezone"`
	TimeAbbreviations string `json:"timezone_abbreviation"`
	HourlyUnits struct{
		Time string `json:"time"`
		Temperature2m string `json:"temperature_2m"`
		Rain string `json:"rain"`
		WindSpeed10m string `json:"wind_speed_10m"`
		WindDirection10m string `json:"wind_direction_10m"`
		WindGusts10m string `json:"wind_gusts_10m"`
		WeatherCode string `json:"weather_code"`
		PressureMsl string `json:"pressure_msl"`
		SoilTemperature0cm string `json:"soil_temperature_0cm"`
	} `json:"hourly_units"`
	Hourly struct{
		Time []string `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
		Rain []float64 `json:"rain"`
		WindSpeed10m []float64 `json:"wind_speed_10m"`
		WindDirection10m []float64 `json:"wind_direction_10m"`
		WindGusts10m []float64 `json:"wind_gusts_10m"`
		WeatherCode []int `json:"weather_code"`
		PressureMsl []float64 `json:"pressure_msl"`
		SoilTemperature0cm []float64 `json:"soil_temperature_0cm"`
	} `json:"hourly"`
	
}
type AirQualityData struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone string `json:"timezone"`
	TimeAbbreviations string `json:"timezone_abbreviation"`
	HourlyUnits struct{
		Time string `json:"time"`
		PM2_5 string `json:"pm2_5"`
		PM10 string `json:"pm10"`
		US_AQI string `json:"us_aqi"`
		EuropeanAQI string `json:"european_aqi"`
	} `json:"hourly_units"`
    Hourly struct {
        PM2_5        []float64 `json:"pm2_5"`
        PM10         []float64 `json:"pm10"`
        US_AQI       []int     `json:"us_aqi"`
        EuropeanAQI  []int     `json:"european_aqi"`
    } `json:"hourly"`
}

func health(w http.ResponseWriter, r *http.Request){
     logger.Println("Health method was invoked");
	 response := map[string]string{"message":"Server is healthy"}
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
	 json.NewEncoder(w).Encode(response)
} 
func getWeather(w http.ResponseWriter, r *http.Request){
	logger.Println("Weather method was invoked.")
	latitude,longitude := r.URL.Query().Get("latitude"),r.URL.Query().Get("longitude")
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=temperature_2m,rain,wind_speed_10m,wind_direction_10m,wind_gusts_10m,weather_code,pressure_msl,soil_temperature_0cm&timezone=auto",latitude,longitude)
	res,err:= http.Get(url)
	if(err!=nil){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError);
		json.NewEncoder(w).Encode(map[string]string{"error":"Error while fetching weather data"})
		logger.Fatalln("Error while fetching weather data",err)
	}
	defer res.Body.Close()
	var WeatherData weatherData
	if err := json.NewDecoder(res.Body).Decode(&WeatherData); err != nil {
		logger.Println("Error while decoding weather data", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error":"Error while decoding weather data"})
		logger.Fatalln("Error while decoding weather data",err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(WeatherData)
}

func getAirQuality(w http.ResponseWriter, r *http.Request){
	logger.Println("Air Quality method was invoked");
	latitude,longitude := r.URL.Query().Get("latitude"),r.URL.Query().Get("longitude")
	url:= fmt.Sprintf("https://air-quality-api.open-meteo.com/v1/air-quality?latitude=%s&longitude=%s&hourly=pm10,pm2_5,european_aqi,us_aqi",latitude,longitude)
	res,err:= http.Get(url);
	if(err!=nil){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Error occured while fetching Air Quality Data "})
		logger.Fatalln("Error occured while fetching Air Quality Data",err);
	}
	defer res.Body.Close();
	var airQualityData AirQualityData;
	jsonErr:=json.NewDecoder(res.Body).Decode(&airQualityData);
	if(jsonErr!=nil){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"Error occured while decoding Air Quality Data "})
		logger.Fatalln("Error occured while decoding Air Quality Data",err);
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airQualityData)
}