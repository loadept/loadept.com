import { ContactSection } from '../modules/about/components/contact-section'
import { InterestsSection } from '../modules/about/components/interests-section'
import { ExperienceSection } from '../modules/about/components/experience-section'
import { IntroSection } from '../modules/about/components/intro-section'

export const About = () => {
  return (
    <div className="container mx-auto px-4 max-w-3xl mb-20">
      <div className="space-y-16">
        <IntroSection />
        <ExperienceSection />
        <InterestsSection />
        <ContactSection />
      </div>
    </div>
  )
}
