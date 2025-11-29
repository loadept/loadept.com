import axios from 'axios'

const apiClient = axios.create({
  baseURL: `${import.meta.env.PUBLIC_API_URL}/api/v1`
})

export default apiClient
