export const Interests = () => {
  return (
    <section className='space-y-3'>
      <div className='flex items-center gap-3'>
        <span className='text-3xl text-[#e06c75]'></span>
        <code className='text-[#c678dd]'>intereses.sh</code>
      </div>

      <div className='text-[#abb2bf] grid grid-cols-1 md:grid-cols-2 gap-4'>
        <div className='flex items-center gap-3 bg-[#282c34] p-4 rounded-lg'>
          <span className='text-3xl text-[#e5c07b]'></span>
          <span className='text-sm md:text-base'>Desarrollo backend</span>
        </div>
        <div className='flex items-center gap-3 bg-[#282c34] p-4 rounded-lg'>
          <span className='text-3xl text-[#61afef]'></span>
          <span className='text-sm md:text-base'>Tecnologías web modernas</span>
        </div>
        <div className='flex items-center gap-3 bg-[#282c34] p-4 rounded-lg'>
          <span className='text-3xl text-[#c678dd]'></span>
          <span className='text-sm md:text-base'>Aprendizaje continuo</span>
        </div>
        <div className='flex items-center gap-3 bg-[#282c34] p-4 rounded-lg'>
          <span className='text-3xl text-[#98c379]'></span>
          <span className='text-sm md:text-base'>Código Abierto</span>
        </div>
      </div>
    </section>
  )
}
