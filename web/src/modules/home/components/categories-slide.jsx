export const CategoriesSlide = ({ categoryData, activeCategory, setActiveCategory }) => {
  return (
    <div className='flex items-center justify-between'>
      <div className='flex items-center space-x-2 overflow-x-auto pb-2 scrollbar-hide'>
        {categoryData.categories.map((category, k) => (
          <button
            key={k}
            onClick={() => setActiveCategory(category.name)}
            className={`flex items-center px-3 py-1.5 rounded-md outline-none transition-colors whitespace-nowrap text-sm
                  ${activeCategory == category.name
                ? `bg-[#282c34]`
                : 'text-[#abb2bf] hover:text-[#528bff]'
              }`}
            style={activeCategory == category.name
              ? { color: category.hex_color }
              : {}
            }
          >
            <span className='text-2xl mr-2'>{category.nerd_icon}</span>
            {category.name}
          </button>
        ))}
      </div>
    </div>
  )
}
