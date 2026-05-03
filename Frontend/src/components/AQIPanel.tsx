import AqiCard from "./AQICard";
import type {AirQualityData,AirQualityCurrent,AirQualityUnits} from '../types';
import { aqiClassifiers } from "../utils";
export default function AqiPanel( aqi : AirQualityData) {

  return (
    <div className="mb-5">
      <div>Air Quality Index</div>
 
      <div className="grid grid-cols-[repeat(auto-fit,minmax(190px,1fr))] gap-3.5">
        {aqi && aqi.current && (Object.keys(aqi.current)as (keyof AirQualityCurrent)[]).map((item:string,index:number)=>{
            {
              let func = aqiClassifiers[item as keyof typeof aqiClassifiers]
              return <AqiCard
                key={index}
                label={item}
                value={aqi.current[item as keyof AirQualityCurrent]}
                unit={aqi.current_units[item as keyof AirQualityUnits]}
                level={func(aqi.current[item as keyof AirQualityCurrent])}
                />}
        })}
      </div>
    </div>
  );
}