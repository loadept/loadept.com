import { useParams } from 'react-router'
import { useEffect, useState } from 'react'
import fm from 'front-matter'
import { fetchArticle } from '../actions/fetchArticle'
import { useNavigate } from 'react-router'

export const useContent = () => {
  const params = useParams()
  const [content, setContent] = useState('')
  const [meta, setMeta] = useState({})
  const navigate = useNavigate()

  const getRaw = async () => {
    const rawMd = await fetchArticle(params)
    if (rawMd.length === 0) {
      navigate('/404')
      return
    }

    const { attributes, body } = fm(rawMd)
    setContent(body)
    setMeta(attributes)
  }

  useEffect(() => {
    getRaw()
  }, [content, params, navigate])

  return { content, meta }
}
