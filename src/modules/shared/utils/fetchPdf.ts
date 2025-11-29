import type { AxiosError } from 'axios'
import apiClient from './apiClient'

export const fetchPdf = async (
  { action, file, quality }:
  { action: string, file: File, quality: string }
) => {
  try {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('quality', quality)

    const res = await apiClient.post(`/pdf/${action}`, formData, {
      responseType: 'arraybuffer',
    })

    return { fileCompressed: res.data }
  } catch (error) {
    const err = error as AxiosError

    console.error(err.response)
    return { file: null }
  }
}
