export const Footer = () => {
  const currentDate = new Date().getFullYear()

  return (
    <footer className='py-5 bg-linear-to-t from-[#282c34] via-90% via-[#282c34] to-[#1f2329]'>
      <div className='container mx-auto px-4'>
        <div className='flex items-center justify-between text-center'>
          <p className='text-xs text-[#abb2bf]'>
            <span className='text-[#c678dd]'>loadept</span> <span className='text-[#e06c75]'>-</span> {currentDate}
          </p>
          <a href='https://github.com/loadept/loadept.com' target='_blank' title='source' className='text-xs'>
            <span className='outline-none text-[#e06c75] border-b'>Código fuente</span>
          </a>
        </div>
      </div>
    </footer>
  )
}
