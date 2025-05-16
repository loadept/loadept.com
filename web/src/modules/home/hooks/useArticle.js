import { useState, useEffect } from 'preact/hooks'
import { fetchArticle } from '../actions/fethArticle'

export const useArticle = (activeCategory) => {
  const [articlesData, setArticlesData] = useState({ articles: [] })

  const getArticles = async () => {
    const articles = await fetchArticle(activeCategory)
    setArticlesData(articles)
  }

  useEffect(() => {
    if (!activeCategory) return

    getArticles()
  }, [activeCategory])

  return { articlesData }
}
