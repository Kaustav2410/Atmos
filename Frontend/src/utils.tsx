import type{ AqiLevel } from "./types";


function classifypm2_5(v: number): AqiLevel {
  if (v <= 12)   return "good";
  if (v <= 35.4) return "moderate";
  if (v <= 55.4) return "unhealthy";
  return "hazardous";
}
 
function classifypm10(v: number): AqiLevel {
  if (v <= 54)  return "good";
  if (v <= 154) return "moderate";
  if (v <= 254) return "unhealthy";
  return "hazardous";
}
 
function classifyeuropeanAqi(v: number): AqiLevel {
  if (v <= 20) return "good";
  if (v <= 60) return "moderate";
  if (v <= 80) return "unhealthy";
  return "hazardous";
}
 
function classifyusAqi(v: number): AqiLevel {
  if (v <= 50)  return "good";
  if (v <= 100) return "moderate";
  if (v <= 200) return "unhealthy";
  return "hazardous";
}
export const aqiClassifiers = {
    "pm2_5" :classifypm2_5, 
    "pm10" :classifypm10, 
    "us_aqi" :classifyusAqi, 
    "european_aqi" :classifyeuropeanAqi, 
}