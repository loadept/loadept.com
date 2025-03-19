import { apiClient } from '../../../core/utils/apiClient'

export const fetchArticle = async ({ category, name }) => {
  try {
    const res = await apiClient.get(`api/article/${category}/${name}`)

    return res.data
  } catch (err) {
    console.error(err.response.data)
    return ''
  }
} 
