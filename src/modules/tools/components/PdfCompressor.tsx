import {
  Upload,
  X,
  Download,
  Loader2,
  ShieldCheck,
  Gauge,
  Zap,
  FileCheck,
  TrendingDown
} from 'lucide-preact'
import { useState, useRef } from 'preact/hooks'
import { fetchPdf } from '../../shared/utils/fetchPdf'
import { FileInput } from './FileInput'

const compressionQualities = [
  {
    id: "low",
    name: "Baja",
    description: "Compresión mínima, máxima calidad",
    icon: ShieldCheck,
    color: "text-[#98c379]",
  },
  {
    id: "normal",
    name: "Normal",
    description: "Balance entre tamaño y calidad",
    icon: Gauge,
    color: "text-[#61afef]",
  },
  {
    id: "extreme",
    name: "Extrema",
    description: "Máxima compresión, menor calidad",
    icon: Zap,
    color: "text-[#e06c75]",
  },
]

export const PdfCompressor = () => {
  const [file, setFile] = useState<File | null>(null)
  const [isProcessing, setIsProcessing] = useState(false)
  const [processedFile, setProcessedFile] = useState<ArrayBuffer | null>(null)
  const [quality, setQuality] = useState("normal")
  const fileInputRef = useRef<HTMLInputElement>(null)

  const handleFileSelect = (e: Event) => {
    const target = e.target as HTMLInputElement
    const selectedFile = target.files?.[0]
    if (selectedFile && selectedFile.type === "application/pdf") {
      setFile(selectedFile)
      setProcessedFile(null)
    }
  }

  const handleDrop = (e: DragEvent) => {
    e.preventDefault()
    const droppedFile = e.dataTransfer?.files[0]
    if (droppedFile && droppedFile.type === "application/pdf") {
      setFile(droppedFile)
      setProcessedFile(null)
    }
  }

  const handleDragOver = (e: DragEvent) => {
    e.preventDefault()
  }

  const removeFile = () => {
    setFile(null)
    setProcessedFile(null)
    setQuality("normal")
    setIsProcessing(false)
    if (fileInputRef.current) {
      fileInputRef.current.value = ""
    }
  }

  const compressPdf = async () => {
    if (!file) return

    setIsProcessing(true)

    const { fileCompressed } = await fetchPdf({ action: 'compress', file, quality })
    if (!fileCompressed) {
      setIsProcessing(false)
      return
    }
    setProcessedFile(fileCompressed)
    setIsProcessing(false)
  }

  const downloadCompressed = () => {
    if (!processedFile || !file) return

    const fileName = `${file.name.split('.')[0]}_compressed.pdf`
    const blob = new Blob([processedFile], { type: 'application/pdf' })
    const url = URL.createObjectURL(blob)
    let a: HTMLAnchorElement | null = document.createElement("a")
    a.href = url
    a.download = fileName
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    URL.revokeObjectURL(url)

    a = null
  }

  return (
    <div className="space-y-20">
      <div className="bg-[#282c34] rounded-lg p-6">
        <h3 className="text-lg font-bold text-[#e5c07b] mb-4">Comprimir PDF</h3>
        {!file ? (
          <FileInput
            handleDrop={handleDrop}
            handleDragOver={handleDragOver}
            handleFileSelect={handleFileSelect}
            fileInputRef={fileInputRef}
          />
        ) : (
          <div className="space-y-4">
            <div className="bg-[#21252b] rounded-lg p-4 flex items-center justify-between">
              <div className="flex items-center">
                <div className="bg-[#e06c75] bg-opacity-20 p-2 rounded mr-3">
                  <Upload className="h-5 w-5 text-[#ffffff]" />
                </div>
                <div>
                  <p className="text-[#e5c07b] font-medium break-all">{file.name}</p>
                  <p className="text-sm text-[#5c6370]">{(file.size / 1024 / 1024).toFixed(2)} MB</p>
                </div>
              </div>
              <button onClick={removeFile} className="text-[#e06c75] hover:text-[#ff6b7a] transition-colors cursor-pointer">
                <X className="h-5 w-5" />
              </button>
            </div>

            {!processedFile && (
              <div className="space-y-4">
                <div className="flex items-center gap-2">
                  <code className="text-[#c678dd]">Calidad de compresión =</code>
                  <code className="text-[#98c379]">"{quality}"</code>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-3">
                  {compressionQualities.map((option) => (
                    <button
                      key={option.id}
                      onClick={() => setQuality(option.id)}
                      className={`p-4 rounded-lg text-left transition-all ${
                        quality === option.id
                          ? "bg-[#21252b] border-2 border-[#528bff]"
                          : "bg-[#21252b] border-2 border-transparent hover:border-[#3e4451]"
                      }`}
                    >
                      <div className="flex items-start mb-2">
                        <option.icon className={`h-5 w-5 mr-2 ${option.color}`} />
                        <div className="flex-1">
                          <h4
                            className={`font-bold ${quality === option.id ? option.color : "text-[#e5c07b]"}`}
                          >
                            {option.name}
                          </h4>
                          <p className="text-sm text-[#abb2bf] mt-1">{option.description}</p>
                        </div>
                      </div>
                    </button>
                  ))}
                </div>
              </div>
            )}

            {!processedFile ? (
              <button
                onClick={compressPdf}
                disabled={isProcessing}
                className="w-full bg-[#21252b] text-[#61afef] px-6 py-3 rounded-md hover:bg-[#2c313a] transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
              >
                {isProcessing ? (
                  <>
                    <Loader2 className="h-5 w-5 mr-2 animate-spin" />
                    Comprimiendo...
                  </>
                ) : (
                  "Comprimir PDF"
                )}
              </button>
            ) : (
              <div className="space-y-4">
                <div className="bg-[#282c34] border border-[#809e6b] rounded-lg p-6 space-y-4">
                  <div className="flex items-center gap-2 mb-4">
                    <FileCheck className="h-6 w-6 text-[#98c379]" />
                    <h4 className="text-lg font-bold text-[#98c379]">Compresión completada</h4>
                  </div>
                  <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                    <div className="bg-[#21252b] rounded-lg p-4 text-center">
                      <p className="text-xs text-[#5c6370] mb-2 uppercase tracking-wide">Tamaño Original</p>
                      <p className="text-2xl font-bold text-[#abb2bf] font-mono">
                        {
                          (file.size / 1024) < 1000
                            ?
                            <>
                              {(file.size / 1024).toFixed(2)}
                              <span className="text-sm text-[#5c6370] ml-1">KB</span>
                            </>
                            :
                            <>
                              {(file.size / 1024 / 1024).toFixed(2)}
                              <span className="text-sm text-[#5c6370] ml-1">MB</span>
                            </>
                        }
                      </p>
                    </div>
                    <div className="bg-[#21252b] rounded-lg p-4 text-center border-2 border-[#98c379]">
                      <p className="text-xs text-[#98c379] mb-2 uppercase tracking-wide font-bold">
                        Tamaño Comprimido
                      </p>
                      <p className="text-2xl font-bold text-[#98c379] font-mono">
                        {
                          (processedFile.byteLength / 1024) < 1000
                            ?
                            <>
                              {(processedFile.byteLength / 1024).toFixed(2)}
                              <span className="text-sm ml-1">KB</span>
                            </>
                            :
                            <>
                              {(processedFile.byteLength / 1024 / 1024).toFixed(2)}
                              <span className="text-sm ml-1">MB</span>
                            </>
                        }
                      </p>
                    </div>
                    <div className="bg-[#21252b] rounded-lg p-4 text-center">
                      <div className="flex items-center justify-center gap-1 mb-2">
                        <TrendingDown className="h-3 w-3 text-[#98c379]" />
                        <p className="text-xs text-[#5c6370] uppercase tracking-wide">Reducción</p>
                      </div>
                      <p className="text-2xl font-bold text-[#98c379] font-mono">
                        {(((file.size - processedFile.byteLength) / file.size) * 100).toFixed(2)}
                        <span className="text-xl">%</span>
                      </p>
                    </div>
                  </div>
                </div>

                <div className="bg-[#98c379] bg-opacity-10 border border-[#98c379] rounded-lg p-4">
                  <p className="text-[#ffffff] text-sm">
                    <code>✓ PDF comprimido exitosamente</code>
                  </p>
                </div>
                <button
                  onClick={downloadCompressed}
                  className="w-full bg-[#21252b] text-[#98c379] px-6 py-3 rounded-md hover:bg-[#2c313a] transition-colors flex items-center justify-center cursor-pointer"
                >
                  <Download className="h-5 w-5 mr-2" />
                  Descargar PDF comprimido
                </button>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  )
}
