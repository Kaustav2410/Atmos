interface WeatherUnits {
  time: string
  temperature_2m: string
  rain: string
  wind_speed_10m: string
  wind_direction_10m: string
  wind_gusts_10m: string
  weather_code: string
  pressure_msl: string
  soilTemperature_0cm: string
}

interface WeatherHourly {
  time: string[]
  temperature_2m: number[]
  rain: number[]
  wind_speed_10m: number[]
  wind_direction_10m: number[]
  wind_gusts_10m: number[]
  weather_code: number[]
  pressure_msl: number[]
  soilTemperature_0cm: number[]
}
export interface WeatherCurrent {
  time: string
  temperature_2m: number
  rain: number
  wind_speed_10m: number
  wind_direction_10m: number
  wind_gusts_10m: number
  weather_code: number 
  pressure_msl: number
  soilTemperature_0cm: number
}
export interface WeatherData {
  latitude: number
  longitude: number
  timezone: string
  timezone_abbreviation: string
  current_units: WeatherUnits
  hourly?: WeatherHourly
  current: WeatherCurrent
}
export interface AirQualityUnits {
  time: string
  pm2_5: string
  pm10: string
  usAqi: string
  europeanAqi: string
}

interface AirQualityHourly {
  pm2_5: number[]
  pm10: number[]
  usAqi: number[]
  europeanAqi: number[]
}
export interface AirQualityCurrent {
  pm2_5: number
  pm10: number
  usAqi: number
  europeanAqi: number
}

export interface AirQualityData {
  latitude: number
  longitude: number
  timezone: string
  timezoneAbbreviation: string
  current_units: AirQualityUnits
  hourly?: AirQualityHourly
  current: AirQualityCurrent
}
export const WMO_DESC: Record<number, string> = {
  0:  "Clear sky",
  1:  "Mainly clear",
  2:  "Partly cloudy",
  3:  "Overcast",
  45: "Fog",
  48: "Icy fog",
  51: "Light drizzle",
  53: "Moderate drizzle",
  55: "Dense drizzle",
  61: "Slight rain",
  63: "Moderate rain",
  65: "Heavy rain",
  71: "Slight snow",
  73: "Moderate snow",
  75: "Heavy snow",
  80: "Slight showers",
  81: "Moderate showers",
  82: "Violent showers",
  95: "Thunderstorm",
  96: "Thunderstorm w/ hail",
};
 
export const WMO_ICON: Record<number, string> = {
  0:  "☀️",
  1:  "🌤️",
  2:  "⛅",
  3:  "☁️",
  45: "🌫️",
  48: "🌫️",
  51: "🌦️",
  53: "🌦️",
  55: "🌧️",
  61: "🌧️",
  63: "🌧️",
  65: "🌧️",
  71: "🌨️",
  73: "❄️",
  75: "❄️",
  80: "🌦️",
  81: "🌦️",
  82: "⛈️",
  95: "⛈️",
  96: "⛈️",
};

export type AqiLevel = "good" | "moderate" | "unhealthy" | "hazardous";

export const LEVEL_COLOR: Record<AqiLevel, string> = {
  good:      "#6fcf6f",
  moderate:  "#e3b715",
  unhealthy: "#e07b39",
  hazardous: "#e05252",
};
 
export const LEVEL_BG: Record<AqiLevel, string> = {
  good:      "rgba(111,207,111,0.10)",
  moderate:  "rgba(227,183,21,0.12)",
  unhealthy: "rgba(224,123,57,0.12)",
  hazardous: "rgba(224,82,82,0.12)",
};
 
export const LEVEL_LABEL: Record<AqiLevel, string> = {
  good:      "Good",
  moderate:  "Moderate",
  unhealthy: "Unhealthy",
  hazardous: "Hazardous",
};
 
// Tailwind arbitrary-value classes per level (border, text)
export const LEVEL_BORDER_CLASS: Record<AqiLevel, string> = {
  good:      "border-l-[#6fcf6f]",
  moderate:  "border-l-[#e3b715]",
  unhealthy: "border-l-[#e07b39]",
  hazardous: "border-l-[#e05252]",
};
 
export const LEVEL_TEXT_CLASS: Record<AqiLevel, string> = {
  good:      "text-[#6fcf6f]",
  moderate:  "text-[#e3b715]",
  unhealthy: "text-[#e07b39]",
  hazardous: "text-[#e05252]",
};