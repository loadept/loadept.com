// import { useParams } from 'react-router'
import { useEffect, useState } from 'preact/hooks'
import fm from 'front-matter'
import { fetchArticle } from '../actions/fetchArticle'
import { route } from 'preact-router'

export const useContent = (params) => {
  // const params = useParams()
  const [content, setContent] = useState('')
  const [meta, setMeta] = useState({})

  const getRaw = async () => {
    const rawMd = await fetchArticle(params)
    if (rawMd.length === 0) {
      route('/404')
      return
    }

    const { attributes, body } = fm(rawMd)
    setContent(body)
    setMeta(attributes)
  }

  useEffect(() => {
    getRaw()
  }, [content])

  return { content, meta }
}
