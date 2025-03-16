import { useParams } from 'react-router'
import { useEffect, useState } from 'react'
import fm from 'front-matter'
import { fetchArticle } from '../actions/fetchArticle'

export const useContent = () => {
  const params = useParams()
  const [content, setContent] = useState('')
  const [meta, setMeta] = useState({})

  useEffect(() => {
    const getRaw = async () => {
      const rawMd = await fetchArticle(params)
      const { attributes, body } = fm(rawMd)
      setContent(body)
      setMeta(attributes)
    }
    getRaw()
  }, [content])

  return { content, meta, params }
}
