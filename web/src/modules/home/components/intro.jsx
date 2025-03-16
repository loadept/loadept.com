import { GlitchLogo } from './glitch-logo'
import { IntroText } from './intro-text'

export const Intro = () => {
  return (
    <section className='grid md:grid-cols-[2fr_3fr] gap-5 md:gap-20'>
      <GlitchLogo />
      <IntroText />
    </section>
  )
}
