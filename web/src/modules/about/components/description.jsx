export const Description = () => {
  return (
    <section className="space-y-3">
      <div className="flex items-center gap-3">
        <span className="text-3xl text-[#56b6c2]"></span>
        <code className="text-[#c678dd]">sobre-mí.go</code>
      </div>
      <div className="space-y-4">
        <h1 className="text-3xl font-bold text-[#e5c07b]">
          Hola, soy <span className="text-[#61afef]">Jesus</span>
        </h1>
        <p className="text-sm md:text-base leading-relaxed">
          Desarrollador con experiencia en arquitectura de software escalable y procesamiento de datos en tiempo
          real. Especializado en backend con <strong className="text-[#61afef] font-black">Go</strong>,
          <strong className="text-[#e5c07b] font-black"> Python</strong> y <strong class="text-[#98c379] font-black">Node.js</strong>,
          microservicios, bases de datos como <strong className="text-[#98c379] font-black">MongoDB </strong>
          y <strong className="text-[#2e93d6] font-black">PostgreSQL</strong>, y tecnologías de streaming como <strong
            className="text-[#e5c07b] font-black">Kafka</strong>.
          Con conocimientos en sistemas de
          recomendación, despliegue en <strong className="text-[#fa7970] font-black">AWS</strong> y optimización de rendimiento.
          Enfocado en crear soluciones
          eficientes y escalables.
        </p>
      </div>
    </section>
  )
}
