import type {AqiLevel} from "../types"
import {LEVEL_COLOR,LEVEL_BG,LEVEL_LABEL} from "../types" 
export default function AqiCard({ label, value, unit, level }: { label: string; value: number; unit: string; level: AqiLevel }) {
  const color = LEVEL_COLOR[level];
  const bg    = LEVEL_BG[level];
 
  return (
    <div className="bg-[#3c3e40] border border-[#bebeb4]/10 rounded-2xl px-5 py-4.5 relative overflow-hidden">
      {/* Bottom accent bar */}
      <div
        className="absolute bottom-0 left-0 right-0 h-[0.75"
        style={{ background: color }}
      /> 
 
      <div className="text-[9px] tracking-[1.5px] uppercase text-[#9a9b97] mb-2">
        {label}
      </div>
 
      <div
        className="font-mono text-[38px] font-bold leading-none tracking-[-1px]"
        style={{ color: color }}
      >
        {value < 100 ? value.toFixed(1) : Math.round(value)}
      </div>
 
      <div className="text-[10px] text-[#9a9b97] mt-1">{unit}</div>
 
      <div
        className="inline-block mt-2 px-2.5 py-0.5 rounded-full text-[11px] font-medium"
        style={{ background: bg, color: color }}
      >
         {LEVEL_LABEL[level]}
      </div>
    </div>
  );
}