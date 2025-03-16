import { useState, useEffect } from 'react'
import { fetchArticle } from '../actions/fethArticle'

export const useArticle = (activeCategory) => {
  const [articlesData, setArticlesData] = useState({ articles: [] })

  useEffect(() => {
    if (!activeCategory) return

    const getArticles = async () => {
      const articles = await fetchArticle(activeCategory)
      setArticlesData(articles)
    }
    getArticles()
  }, [activeCategory])

  return { articlesData }
}
