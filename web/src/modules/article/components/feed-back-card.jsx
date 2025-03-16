export const FeedBackCard = ({ repo_url }) => {
  return (
    <div className='mt-20 border-t border-[#3e4451] pt-8'>
      <div className='bg-[#282c34] rounded-lg p-5'>
        <h3 className='text-md font-bold text-[#e5c07b] mb-4 flex items-center'>
          <span className='text-2xl mr-2 text-[#98c379]'></span>
          ¿Encontraste un error o tienes una sugerencia?
        </h3>

        <p className='text-[#abb2bf] text-sm mb-4'>
          Este artículo es de código abierto. ¿Ves algo incorrecto o que podría mejorarse? Todas las
          correcciones son bienvenidas.
        </p>

        <div className='flex flex-wrap'>
          <a
            href={repo_url}
            target='_blank'
            rel='noopener noreferrer'
            className='inline-flex items-center px-3 py-1 bg-[#21252b] text-[#61afef] rounded-md hover:bg-[#2c313a] transition-colors'
          >
            <span className='text-xl mr-2 text-[#61afef]'></span>
            Editar en GitHub
          </a>
        </div>

        <div className='mt-4 text-sm text-[#5c6370]'>
          <code>// Gracias por ayudar a mejorar este contenido</code>
        </div>
      </div>
    </div>
  )
}
