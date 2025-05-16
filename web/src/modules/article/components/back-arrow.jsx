import { Link } from 'preact-router/match'

export const BackArrow = () => {
  return (
    <Link
      href='/'
      className='inline-flex items-center text-[#61afef] mb-8 hover:text-[#528bff] transition-colors'
    >
      <span className='text-2xl mr-1'></span>
      Volver al inicio
    </Link>
  )
}
