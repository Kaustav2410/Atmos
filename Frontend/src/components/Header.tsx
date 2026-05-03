// components/Header.tsx

export default function Header() {
  return (
    <div className="flex items-end justify-between mb-7 flex-wrap gap-3">
      {/* Brand */}
      <div className="flex items-center gap-3">
        <div
          className="w-10 h-10 rounded-xl flex items-center justify-center text-xl"
          style={{ background: "linear-gradient(135deg, #e3b715, #c8960f)" }}
        >
          🌐
        </div>
        <div>
          <div className="text-[22px] font-bold tracking-tight text-[#f0efe8]">
            AeroSense
          </div>
          <div className="text-[10px] tracking-[2.5px] uppercase text-[#9a9b97] mt-0.5">
            Weather &amp; Air Quality
          </div>
        </div>
      </div>

      {/* Live badge */}
      <div className="flex items-center gap-1.5 bg-[#6fcf6f]/10 border border-[#6fcf6f]/20 rounded-full px-3.5 py-1 text-xs text-[#6fcf6f]">
        <span
          className="w-1.5 h-1.5 rounded-full bg-[#6fcf6f] inline-block"
          style={{ animation: "blink 2s infinite" }}
        />
        Live Data
      </div>
    </div>
  );
}