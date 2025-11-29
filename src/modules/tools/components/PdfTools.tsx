import { useState } from 'preact/hooks';
import { PdfCompressor } from './PdfCompressor';
import { Minimize2, Combine } from 'lucide-preact';

export const PdfTools = () => {
  const [mode, setMode] = useState<'compress' | 'merge'>('compress');

  return (
    <div className="space-y-8">
      <div className="flex items-center space-x-2 bg-[#282c34] p-2 rounded-lg w-fit">
        <button 
          onClick={() => setMode('compress')}
          className={`flex items-center px-3 py-2 rounded-md transition-colors ${
            mode === 'compress' 
              ? 'bg-[#21252b] text-[#61afef]' 
              : 'text-[#abb2bf] hover:text-[#528bff]'
          }`}
        >
          <Minimize2 className="h-4 w-4 mr-2" />
          Comprimir PDF
        </button>
        <button 
          onClick={() => setMode('merge')}
          className={`flex items-center px-3 py-2 rounded-md transition-colors ${
            mode === 'merge' 
              ? 'bg-[#21252b] text-[#61afef]' 
              : 'text-[#abb2bf] hover:text-[#528bff]'
          }`}
        >
          <Combine className="h-4 w-4 mr-2" />
          Unir PDFs
        </button>
      </div>
      <PdfCompressor />
      {/* <div class="mt-8">{mode === "compress" ?  : <PdfMerger />}</div> */}
    </div>
  )
}