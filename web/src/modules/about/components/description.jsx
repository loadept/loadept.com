export const Description = () => {
  const birthdate = new Date('2003-10-01T00:00:00')
  const currentDate = new Date()

  let age = currentDate.getFullYear() - birthdate.getFullYear()

  const hasBirthdayPassed =
    currentDate.getMonth() > birthdate.getMonth() ||
    (currentDate.getMonth() === birthdate.getMonth() &&
    currentDate.getDate() >= birthdate.getDate());

  if (!hasBirthdayPassed) {
    age--;
  }

  const Peru = () => (
    <strong className='text-[#61afef] font-black'>
      <span className='text-red-400 font-black'>P</span><span className='text-gray-300 font-black'>er</span><span className='text-red-400 font-black'>ú</span>
    </strong>
  )

  return (
    <section className='space-y-3'>
      <div className='flex items-center gap-3'>
        <span className='text-3xl text-[#56b6c2]'></span>
        <code className='text-[#c678dd]'>sobre-mí.go</code>
      </div>
      <div className='space-y-4'>
        <h1 className='text-3xl font-bold text-[#e5c07b]'>
          Hola, soy <span className='text-[#61afef]'>Jesus</span>
        </h1>
        <p className='text-sm md:text-base leading-relaxed'>
          Tengo { age } años y vivo en <Peru />. Actualmente trabajo como desarrollador de software en
          la <span className='underline'>Cooperativa de Ahorro y Crédito Fondesurco</span>. Soy un entusiasta
          del software libre y el open source. Me apasiona la tecnología, la programación y la música,
          en especial el <strong className='text-[#c678dd] font-black'>Metal</strong>, que disfruto escuchar mientras programo porque me motiva mucho.
        </p>
        <p className='text-sm md:text-base leading-relaxed'>
          Uno de mis lenguajes favoritos es <strong className='text-[#61afef] font-black'>Go</strong>,
          por su simplicidad y lo divertido que resulta trabajar con él.
          Dentro de mi stack también manejo con solidez <strong className='text-[#e5c07b] font-black'>Python</strong> y
          <strong className='text-[#98c379] font-black'> Node.js</strong>. En mi tiempo libre me gusta explorar nuevas tecnologías,
          escribir notas sobre lo que aprendo y disfrutar de series o películas.
        </p>
      </div>
    </section>
  )
}
