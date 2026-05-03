import type { WeatherData } from "../types";
import {WMO_DESC,WMO_ICON} from "../types"
import StatChip from "../statCard";

export default function WeatherHero(weather:WeatherData | null) {
    if(weather=== null){
        return (
            <div>
                No Data Found
            </div>
        )
    }
  return (
    <div className="bg-[#3c3e40] border border-[#bebeb4]/10 rounded-[18px] p-6 mb-5 grid grid-cols-2 gap-6">
      {/* Temperature side */}
      <div className="flex flex-col justify-center">
        <div className="text-[80px] font-bold leading-none tracking-[-4px] text-[#f0efe8]">
          {Math.round(weather.current.temperature_2m)}
          <span className="text-[36px] text-[#9a9b97]">°C</span>
        </div>
 
        <div className="flex items-center gap-2.5 mt-2.5 text-[#9a9b97] text-sm">
          <span className="text-[28px]">{WMO_ICON[weather.current.weather_code] ?? "🌡️"}</span>
          {WMO_DESC[weather.current.weather_code] ?? "Unknown"}
        </div>
 
        <div className="mt-4 font-mono text-xs text-[#9a9b97]">
        {new Date(weather.current.time).toLocaleTimeString([], { 
            hour: '2-digit', 
            minute: '2-digit', 
            hour12: true 
        })}
        </div>
      </div>
 
      {/* Stat grid */} 
      <div className="grid grid-cols-2 gap-2.5">
        <StatChip label="Rain"        value={weather.current.rain.toFixed(1)}            unit={weather.current_units.rain}   /> 
        <StatChip label="Wind Speed"  value={Math.round(weather.current.wind_speed_10m)} unit={weather.current_units.wind_speed_10m} /> 
        <StatChip label="Wind Gusts"  value={Math.round(weather.current.wind_gusts_10m)} unit={weather.current_units.wind_gusts_10m} />
        <StatChip label="Pressure MSL" value={Math.round(weather.current.pressure_msl)}  unit={weather.current_units.pressure_msl}  />
      </div>
    </div>
  );
}