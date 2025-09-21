import formatDate from '../../../core/utils/format-date'
import { useContext } from 'preact/hooks'
import { DataContext } from '../../../core/providers/dataContext'
import { route } from 'preact-router'

export const ArticlesSection = ({ articlesData, categoryData, activeCategory }) => {
  const { setData } = useContext(DataContext)

  const handleClick = (article, url) => {
    setData(url)
    route(`/articles/${activeCategory}/${article}`)
  }

  return (
    <>
      <div>
        <h3 className='text-xl font-bold text-[#e5c07b] mb-6'>
          {categoryData.categories.find((c) => c.name === activeCategory)?.name}
        </h3>
      </div>

      <div className='space-y-4'>
        {articlesData.articles.length > 0 ? (
          articlesData.articles.map((article, k) => (
            <button
              key={k}
              onClick={() => handleClick(article.name, article.html_url)}
              className='w-full outline-none group flex items-center justify-between
                        py-2 transition-colors hover:bg-[#282c34] rounded-md px-3'
            >
              <div className='flex items-center space-x-3'>
                <div className={`h-2 w-2 rounded-full `} />
                <span className='text-[#abb2bf] group-hover:text-[#e5c07b] transition-colors'>{article.name}</span>
              </div>
              <span className='text-sm text-wrap text-[#5c6370]'>Actualizado: {formatDate(article.updated_at, 'long')}</span>
            </button>
          ))
        ) : (
          <div className='text-center py-4 text-[#5c6370]'>No hay artículos en esta categoría</div>
        )}
      </div>
    </>
  )
}
