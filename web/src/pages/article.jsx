import { useContext } from 'react'
import { BackArrow } from '../modules/article/components/back-arrow'
import { MainContent } from '../modules/article/components/main-content'
import { useContent } from '../modules/article/hooks/useContent'
import { DataContext } from '../core/providers/context'

export const Article = () => {
  const { content, meta } = useContent()
  const { data } = useContext(DataContext)

  return (
    <div className='container mx-auto px-4 max-w-3xl mb-20'>
      <BackArrow />
      <MainContent content={content} meta={meta} repo_url={data} />
    </div>
  )
}
