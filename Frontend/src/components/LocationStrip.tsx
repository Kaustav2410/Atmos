import type { WeatherData } from "../types";
export default function LocationStrip({ data }: { data: WeatherData }) {
  const fields: [string, string][] = [
    ["Latitude",     `${data.latitude}°`],
    ["Longitude",    `${data.longitude}°`],
    ["Timezone",     data.timezone],
    ["Abbreviation", data.timezone_abbreviation],
  ];
 
  return (
    <div className="bg-[#3c3e40] border border-[#bebeb4]/10 rounded-[14px] px-5 py-3.5 mb-5 grid grid-cols-[repeat(auto-fit,minmax(540px,4fr))] gap-4">
      {fields.map(([label, val]) => (
        <div key={label}>
          <div className="text-[9px] tracking-[1.4px] uppercase text-[#9a9b97] mb-1">
            {label}
          </div>
          <div className="font-mono text-[12.5px] text-[#e3b715]">{val}</div>
        </div>
      ))}
    </div>
  );
}