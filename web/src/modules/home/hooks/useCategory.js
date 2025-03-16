import { useEffect, useState } from 'react'
import { fetchCategorie } from '../actions/fetchCategory'

export const useCategory = () => {
  const [activeCategory, setActiveCategory] = useState(null)
  const [categoryData, setCategoryData] = useState({ categories: [] })

  useEffect(() => {
    const getCategories = async () => {
      const categories = await fetchCategorie()
      setCategoryData(categories)
      setActiveCategory(categories.categories[0].name)
    }
    getCategories()
  }, [])

  return { categoryData, activeCategory, setCategoryData, setActiveCategory }
}
