interface Props {
  label: string;
  value: string | number;
  unit?: string;
}
 
export default function StatChip({ label, value, unit }: Props) {
  return (
    <div className="bg-[#404244] border border-[#bebeb4]/10 rounded-xl px-3.5 py-2.5">
      <div className="text-[9px] tracking-[1.4px] uppercase text-[#9a9b97] mb-1.5">
        {label}
      </div>
      <div className="font-mono text-[17px] font-medium text-[#f0efe8]">
        {value}
        {unit && (
          <span className="text-[11px] text-[#9a9b97] ml-1">{unit}</span>
        )}
      </div>
    </div>
  );
}