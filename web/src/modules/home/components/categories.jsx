import { ArticlesSection } from './articles-section'
import { CategoriesSlide } from './categories-slide'
import { useCategory } from '../hooks/useCategory'
import { useArticle } from '../hooks/useArticle'

export const Categories = () => {
  const {
    categoryData,
    activeCategory,
    setActiveCategory
  } = useCategory()

  const { articlesData } = useArticle(activeCategory)

  return (
    <section className='space-y-6 mt-15'>
      <div className='flex items-center gap-3'>
        <span className='text-3xl text-[#98c379]'></span>
        <h2 className='text-2xl font-bold text-[#e5c07b]'>Contenido</h2>
      </div>

      <div className='space-y-6'>
        <CategoriesSlide
          categoryData={categoryData}
          activeCategory={activeCategory}
          setActiveCategory={setActiveCategory}
        />
        <ArticlesSection
          articlesData={articlesData}
          categoryData={categoryData}
          activeCategory={activeCategory}
        />
      </div>
    </section>
  )
}
