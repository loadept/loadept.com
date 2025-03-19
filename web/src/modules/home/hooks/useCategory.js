import { useEffect, useState } from 'react'
import { fetchCategorie } from '../actions/fetchCategory'

export const useCategory = () => {
  const [activeCategory, setActiveCategory] = useState(null)
  const [categoryData, setCategoryData] = useState({ categories: [] })


  const getCategories = async () => {
    const categories = await fetchCategorie()

    setCategoryData(categories)
    setActiveCategory(categories.categories[0].name)
  }

  useEffect(() => {
    getCategories()
  }, [])

  return { categoryData, activeCategory, setCategoryData, setActiveCategory }
}
