import { CategoriesSection } from "./components/categories-section"

export const Index = () => {
  return (
    <div className="container mx-auto px-4">
      <section id="content" className="grid md:grid-cols-[2fr_3fr] gap-5 md:gap-20">
        <div className="flex flex-col items-center md:justify-self-end justify-self-center">
          <div id="logo" className="md:size-[18rem] rounded-full size-[13rem]">
          </div>
        </div>
        <article
          className="self-center md:justify-self-start justify-self-center text-sm md:text-base md:text-left text-center">
          <p>
            Aquí comparto lo que aprendo sobre <strong className="text-[#c678dd]">Linux</strong>,
            <strong className="text-[#e06c75]">programación</strong>,
          </p>
          <p>
            <strong className="text-[#98c379]">tecnología</strong>,
            <strong className="text-[#61afef]">Go</strong> y más,
            con proyectos y notas en Markdown.
          </p>
          <p>
            Espero que te sirva tanto como a mí. <span className="text-amber-300 text-3xl">󰈸</span>
          </p>
        </article>
      </section>
      <CategoriesSection />
    </div>
  )
}
