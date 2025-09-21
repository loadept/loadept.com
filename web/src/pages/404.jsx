import { Link } from 'preact-router/match'
import { useEffect } from 'preact/hooks'

export const NotFound = () => {
  useEffect(() => {
    document.title = '404 - Página no encontrada - loadept'
  }, [])

  return (
    <div className='container px-4 py-12 max-w-lg mx-auto'>
      <div className='space-y-6 text-center'>
        <div className='flex items-center justify-center gap-3'>
          <span className='text-2xl text-[#e06c75]'></span>
          <code className='text-[#c678dd]'>404</code>
        </div>

        <div className='space-y-4'>
          <pre className='text-6xl font-bold tracking-wider text-[#e06c75]'>
            <code>404</code>
          </pre>
          <h1 className='text-xl text-[#abb2bf]'>
            <span className='text-[#e5c07b]'>Error:</span> Página no encontrada
          </h1>
          <p className='text-[#5c6370] text-sm'>
            <code>La URL solicitada no se encontró en este servidor.</code>
          </p>
        </div>

        <div className='pt-6 hover:text-[#528bff]'>
          <Link
            href='/'
            className='inline-flex items-center gap-2 text-lg text-[#61afef] transition-colors animate-pulse'
          >
            <span className='text-[#98c379]'>$</span> cd /
          </Link>
        </div>
      </div>
    </div>
  )
}
