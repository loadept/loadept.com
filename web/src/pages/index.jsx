import { useEffect } from 'preact/hooks'
import { Categories } from '../modules/home/components/categories'
import { Intro } from '../modules/home/components/intro'

export const Index = () => {
  useEffect(() => {
    document.title = 'Inicio - loadept'
  }, [])

  return (
    <div className='container mx-auto px-4'>
      <Intro />
      <Categories />
    </div>
  )
}
