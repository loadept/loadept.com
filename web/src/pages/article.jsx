import { useContext, useEffect } from 'preact/hooks'
import { BackArrow } from '../modules/article/components/back-arrow'
import { MainContent } from '../modules/article/components/main-content'
import { useContent } from '../modules/article/hooks/useContent'
import { DataContext } from '../core/providers/dataContext'

export const Article = ({ category, name }) => {
  const { content, meta } = useContent({ category, name })
  const { data } = useContext(DataContext)

  useEffect(() => {
    document.title = `${name} - loadept`
  }, [])

  return (
    <div className='container mx-auto px-4 max-w-3xl mb-20'>
      <BackArrow />
      <MainContent content={content} meta={meta} repo_url={data} />
    </div>
  )
}
