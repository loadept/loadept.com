import { ContactSection } from '../modules/about/components/contact-section'
import { Interests } from '../modules/about/components/interests'
import { Experience } from '../modules/about/components/experience'
import { Description } from '../modules/about/components/description'
import { useEffect } from 'preact/hooks'

export const About = () => {
  useEffect(() => {
    document.title = 'Sobre mi - loadept'
  }, [])

  return (
    <div className='container mx-auto px-4 max-w-3xl mb-20'>
      <div className='space-y-16'>
        <Description />
        <Experience />
        <Interests />
        <ContactSection />
      </div>
    </div>
  )
}
