package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
	"strconv"
)
var logger = log.Default()
var httpClient = &http.Client{
	Timeout: 10*time.Second,
}
type WeatherUnits struct{
		Time string `json:"time"`
		Temperature2m string `json:"temperature_2m"`
		Rain string `json:"rain"`
		WindSpeed10m string `json:"wind_speed_10m"`
		WindDirection10m string `json:"wind_direction_10m"`
		WindGusts10m string `json:"wind_gusts_10m"`
		WeatherCode string `json:"weather_code"`
		PressureMsl string `json:"pressure_msl"`
		SoilTemperature0cm string `json:"soil_temperature_0cm"`
	} 
type weatherData struct{
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone string `json:"timezone"`
	TimeAbbreviations string `json:"timezone_abbreviation"`
	Units WeatherUnits `json:"current_units"`
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
	Current struct{
		Time string `json:"time"`
		Temperature2m float64 `json:"temperature_2m"`
		Rain float64 `json:"rain"`
		WindSpeed10m float64 `json:"wind_speed_10m"`
		WindDirection10m float64 `json:"wind_direction_10m"`
		WindGusts10m float64 `json:"wind_gusts_10m"`
		WeatherCode int `json:"weather_code"`
		PressureMsl float64 `json:"pressure_msl"`
		SoilTemperature0cm float64 `json:"soil_temperature_0cm"`
	} `json:"current"`
	
}
type airQualityUnits struct{
		Time string `json:"time"`
		PM2_5 string `json:"pm2_5"`
		PM10 string `json:"pm10"`
		US_AQI string `json:"us_aqi"`
		EuropeanAQI string `json:"european_aqi"`
	}
type airQualityData struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone string `json:"timezone"`
	TimeAbbreviations string `json:"timezone_abbreviation"`
	Units airQualityUnits `json:"current_units"`
    Hourly struct {
        PM2_5        []float64 `json:"pm2_5"`
        PM10         []float64 `json:"pm10"`
        US_AQI       []int     `json:"us_aqi"`
        EuropeanAQI  []int     `json:"european_aqi"`
    } `json:"hourly"`
	Current struct {
        PM2_5        float64 `json:"pm2_5"`
        PM10         float64 `json:"pm10"`
        US_AQI       int     `json:"us_aqi"`
        EuropeanAQI  int     `json:"european_aqi"`
    } `json:"current"`
}
type coordinates struct{
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
func health(w http.ResponseWriter, r *http.Request){
     logger.Println("Health method was invoked");
	 response := map[string]string{"message":"Server is healthy"}
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
	 json.NewEncoder(w).Encode(response)
} 
func getCoordinates(ip string) (latitude string,longitude string,errorMessage error){
		logger.Println("No latitude and longitude provided. Fetching location data using IP address:", ip)
		if strings.Contains(ip,"::1") || strings.HasPrefix(ip, "172."){
			logger.Println("Locally hosted therefore can't fetch user Ip please enable location.Using default coordinates for Sydney, Australia")
			return "-33.83","151.144",nil
		} 
		locationUrl := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,countryCode,lat,lon,timezone,as", ip)
		res, err := httpClient.Get(locationUrl)
		if err!=nil{ 
			 return "","",fmt.Errorf("error while fetching the location for IP %s: %w", ip, err)
		}
		defer res.Body.Close()
		var coords coordinates;
		iperr:= json.NewDecoder(res.Body).Decode(&coords);
		if iperr!=nil {
			 return "","",fmt.Errorf("error while decoding the location for IP %s: %w", ip, iperr)
		}
		logger.Printf("Fetched location data using IP address: Latitude: %f, Longitude: %f", coords.Latitude, coords.Longitude)
		 return strconv.FormatFloat(coords.Latitude, 'f', -1, 64),strconv.FormatFloat(coords.Longitude, 'f', -1, 64),nil 
		}
func getIpAddress(r *http.Request) (ip string,errorMessage error){
		IP := r.Header.Get("X-Forwarded-For")
		if IP == ""{
			IP = r.RemoteAddr
			host,_,err:=net.SplitHostPort(IP);
			if err != nil {
				return "",fmt.Errorf("error while parsing IP address");
			}
			logger.Println("Client IP address from RemoteAddr header:", IP)
			return host,nil;
		}
		IP = strings.Split(r.Header.Get("X-Forwarded-For"),",")[0]
		logger.Println("Client IP address from X-Forwarded-For header:", IP)
		return IP,nil
}
func getWeather(w http.ResponseWriter, r *http.Request){
	latitude,longitude := r.URL.Query().Get("latitude"),r.URL.Query().Get("longitude")
	if latitude == "" || longitude == ""{ 
		ip,ipErrorMessage:= getIpAddress(r);
		if ipErrorMessage != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error":ipErrorMessage.Error()})
				logger.Println(ipErrorMessage.Error())
				return;
			}
		var errorMessage error
		latitude,longitude,errorMessage = getCoordinates(ip);
		if errorMessage!= nil{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error":errorMessage.Error()})
			logger.Println("error while fetching location data",errorMessage)
			return;
		}
	}
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current=temperature_2m,apparent_temperature,relative_humidity_2m,is_day,wind_speed_10m,wind_direction_10m,wind_gusts_10m,precipitation,rain,showers,snowfall,weather_code,cloud_cover,surface_pressure,pressure_msl&timezone=auto",latitude,longitude)
	// https://api.open-meteo.com/v1/forecast?latitude=latitude&longitude=longitude&hourly=temperature_2m,rain,wind_speed_10m,wind_direction_10m,wind_gusts_10m,weather_code,pressure_msl,soil_temperature_0cm&current=temperature_2m,apparent_temperature,relative_humidity_2m,is_day,wind_speed_10m,wind_direction_10m,wind_gusts_10m,precipitation,rain,showers,snowfall,weather_code,cloud_cover,surface_pressure,pressure_msl&timezone=auto
	res,err:= httpClient.Get(url)
	if err!=nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError);
		json.NewEncoder(w).Encode(map[string]string{"error":"error while fetching weather data"})
		logger.Println("error while fetching weather data",err)
		return;
	}
	defer res.Body.Close()
	var WeatherData weatherData
	if err := json.NewDecoder(res.Body).Decode(&WeatherData); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error":"error while decoding weather data"})
		logger.Println("error while decoding weather data",err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(WeatherData)
}

func getAirQuality(w http.ResponseWriter, r *http.Request){
	logger.Println("Air Quality method was invoked");
	latitude,longitude := r.URL.Query().Get("latitude"),r.URL.Query().Get("longitude")
	if latitude == "" || longitude == "" {
		ip,ipErrorMessage:= getIpAddress(r);
		if ipErrorMessage != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error":ipErrorMessage.Error()})
				logger.Println(ipErrorMessage.Error())
				return;
			}
		var errorMessage error
		latitude,longitude,errorMessage = getCoordinates(ip);
		if errorMessage!= nil{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error":errorMessage.Error()})
			logger.Println("error while fetching location data",errorMessage)
			return;
		}
	}
	url:= fmt.Sprintf("https://air-quality-api.open-meteo.com/v1/air-quality?latitude=%s&longitude=%s&current=european_aqi,us_aqi,pm10,pm2_5,carbon_monoxide,nitrogen_dioxide,sulphur_dioxide,ozone",latitude,longitude)
	// https://air-quality-api.open-meteo.com/v1/air-quality?latitude=52.52&longitude=13.41&hourly=pm10,pm2_5,european_aqi,us_aqi&current=european_aqi,us_aqi,pm10,pm2_5,carbon_monoxide,nitrogen_dioxide,sulphur_dioxide,ozone
	res,err:= httpClient.Get(url);
	if err!=nil { 
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"error occured while fetching Air Quality Data "})
		logger.Println("error occured while fetching Air Quality Data",err);
		return;
	}
	defer res.Body.Close();
	var airQualityData airQualityData;
	jsonErr:=json.NewDecoder(res.Body).Decode(&airQualityData);
	if jsonErr!=nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message":"error occured while decoding Air Quality Data "})
		logger.Println("error occured while decoding Air Quality Data",jsonErr);
		return;
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airQualityData)
}