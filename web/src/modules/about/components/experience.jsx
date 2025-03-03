export const Experience = () => {
  return (
    <section className="space-y-3">
      <div className="flex items-center gap-2">
        <span className="text-3xl text-[#98c379]"></span>
        <code className="text-[#c678dd]">experiencia.md</code>
      </div>
      <div className="space-y-6">
        <div className="space-y-4">
          <div className="flex items-center gap-4"> <span class="h-2 w-2 rounded-full bg-[#98c379]"></span>
            <h2 className="text-lg font-bold text-[#e5c07b]">Desarrollador Full Stack @ 10x Tecnología</h2>
          </div>
          <p className="text-[#abb2bf] text-sm md:text-base pl-6">
            Realicé mantenimiento y mejoras en Laravel y Next.js, optimizando rendimiento y funcionalidad.
            Refactoricé módulos para facilitar su mantenimiento, corregí bugs en el ERP y ajusté funcionalidades
            según las necesidades del cliente.
          </p>
        </div>

        <div className="space-y-4">
          <div className="flex items-center gap-4">
            <span className="h-2 w-2 rounded-full bg-[#61afef]"></span>
            <h2 className="text-lg font-bold text-[#e5c07b]">Desarrollador Backend @ Tiksup</h2>
          </div>
          <p className="text-[#abb2bf] text-sm md:text-base pl-6">
            Participación en el desarrollo de Tiksup, una plataforma de videos con recomendaciones en tiempo real,
            utilizando microservicios con Go, Python, Node.js, Kafka, MongoDB, Redis y Spark. Gestión del
            almacenamiento en AWS S3 y despliegue en AWS EC2.
          </p>
        </div>

        <div className="space-y-4">
          <div className="flex items-center gap-4">
            <span className="h-2 w-2 rounded-full bg-[#e06c75]"></span>
            <h2 className="text-lg font-bold text-[#e5c07b]">Desarrollador Full Stack @ Luchadores</h2>
          </div>
          <p className="text-[#abb2bf] text-sm md:text-base pl-6">
            Desarrollo de una aplicación multiplataforma para pacientes con cáncer, con arquitectura de
            microservicios. Implementación del backend con Node.js y Nginx, desplegado en AWS EC2 con S3. Uso de
            Docker para contenerización y entornos consistentes.
          </p>
        </div>
      </div>
    </section>
  )
}
