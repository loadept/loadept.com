export const Experience = () => {
  const experiences = [
    {
      role: 'Analista de Desarrollo @ Coopac Fondesurco',
      description: `Me desempeñé como desarrollador de software en la cooperativa, donde resolví incidencias en el
      sistema legacy y desarrollé nuevas funcionalidades clave para el negocio. Integré la central de riesgos con
      servicios externos, conectando la API principal en Node.js con el sistema legacy en PHP. Además, participé
      en la migración de reportes críticos y en la configuración de servidores PostgreSQL, contribuyendo a mejorar
      la eficiencia operativa y la estabilidad del sistema.`,
      color: '#c678dd'
    },
    {
      role: 'Desarrollador Full Stack @ 10x Tecnología',
      description: `Realicé mantenimiento y mejoras en Laravel y Next.js, optimizando rendimiento y funcionalidad.
      Refactoricé módulos para facilitar su mantenimiento, corregí bugs en el ERP y ajusté funcionalidades
      según las necesidades del cliente.`,
      color: '#98c379'
    },
    {
      role: 'Desarrollador Backend @ Tiksup (Proyecto de Instuto)',
      description: `Participación en el desarrollo de Tiksup, una plataforma de videos con recomendaciones en tiempo real,
      utilizando microservicios con Go, Python, Node.js, Kafka, MongoDB, Redis y Spark. Gestión del
      almacenamiento en AWS S3 y despliegue en AWS EC2.`,
      color: '#61afef'
    },
    {
      role: 'Desarrollador Full Stack @ Luchadores (Proyecto de Instuto)',
      description: `Desarrollo de una aplicación multiplataforma para pacientes con cáncer, con arquitectura de
      microservicios. Implementación del backend con Node.js y Nginx, desplegado en AWS EC2 con S3. Uso de
      Docker para contenerización y entornos consistentes.`,
      color: '#e06c75'
    }
  ]

  return (
    <section className='space-y-3'>
      <div className='flex items-center gap-2'>
        <span className='text-3xl text-[#98c379]'></span>
        <code className='text-[#c678dd]'>experiencia.md</code>
      </div>
      <div className='space-y-6'>
        {
          experiences.map((exp, index) => (
            <div key={index} className='space-y-4'>
              <div className='flex items-center gap-4'>
                <span className={`h-2 w-2 rounded-full`} style={{ backgroundColor: exp.color }}></span>
                <h2 className='text-lg font-bold text-[#e5c07b]'>{exp.role}</h2>
              </div>
              <p className='text-[#abb2bf] text-sm md:text-base pl-6'>
                {exp.description}
              </p>
            </div>
          ))
        }
      </div>
    </section>
  )
}
