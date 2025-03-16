import { Link } from "react-router"
import formatDate from "../../../core/utils/format-date"

export const ArticlesSection = ({ articlesData, categoryData, activeCategory }) => {
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
            <Link
              key={k}
              className='group flex items-center justify-between py-2 transition-colors hover:bg-[#282c34] rounded-md px-3'
              to={`/article/${activeCategory}/${article.name}`}
            >
              <div className='flex items-center space-x-3'>
                <div className={`h-2 w-2 rounded-full `} />
                <span className='text-[#abb2bf] group-hover:text-[#e5c07b] transition-colors'>{article.name}</span>
              </div>
              <span className='text-sm text-wrap text-[#5c6370]'>Actualizado: {formatDate(article.updated_at, 'long')}</span>
            </Link>
          ))
        ) : (
          <div className='text-center py-4 text-[#5c6370]'>No hay artículos en esta categoría</div>
        )}
      </div>
    </>
  )
}
