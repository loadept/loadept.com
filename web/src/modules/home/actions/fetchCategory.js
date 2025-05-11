import { apiClient } from '../../../core/utils/apiClient'

export const fetchCategorie = async () => {
  try {
    const res = await apiClient.get(`api/categories`)

    return res.data
  } catch (err) {
    console.error(err.response.data)
  }
}
