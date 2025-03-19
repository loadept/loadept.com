import { apiClient } from "../../../core/utils/apiClient"

export const fetchArticle = async (articleName) => {
  try {
    const res = await apiClient.get(`api/article/${encodeURIComponent(articleName)}`)

    const jsonData = await res.data
    const processData = {
      ...jsonData,
      articles: jsonData.articles.map(artl => ({
        ...artl,
        name: artl.name.split('.')[0]
      }))

    }

    return processData
  } catch (err) {
    console.error(err.response.data)
  }
}
