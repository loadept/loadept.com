import { Upload } from 'lucide-preact'
import type { RefObject } from 'preact'

export const FileInput = (
  { handleDrop, handleDragOver, handleFileSelect, fileInputRef }:
  { handleDrop: (e: DragEvent) => void,
    handleDragOver: (e: DragEvent) => void,
    handleFileSelect: (e: Event) => void,
    fileInputRef: RefObject<HTMLInputElement>
  }
) => {
  return (
    <div
      onDrop={handleDrop}
      onDragOver={handleDragOver}
      onClick={() => fileInputRef.current?.click()}
      className="border-2 border-dashed border-[#3e4451] rounded-lg p-12 text-center cursor-pointer hover:border-[#528bff] transition-colors"
    >
      <Upload className="h-12 w-12 mx-auto mb-4 text-[#5c6370]" />
      <p className="text-[#abb2bf] mb-2">Arrastra un archivo PDF aquí o haz clic para seleccionar</p>
      <p className="text-sm text-[#5c6370]">
        <code>// Solo archivos PDF</code>
      </p>
      <input
        ref={fileInputRef}
        type="file"
        accept="application/pdf"
        onChange={handleFileSelect}
        className="hidden"
        />
    </div>
  )
}
