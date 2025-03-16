export const fetchArticle = async (articleName) => {
  try {
    const req = await fetch(`${API_URL}api/article/${encodeURIComponent(articleName)}`)
    if (req.status != 200) {
      throw new Error('Error to obtain data')
    }

    const jsonData = await req.json()
    const processData = {
      ...jsonData,
      articles: jsonData.articles.map(artl => ({
        ...artl,
        name: artl.name.split('.')[0]
      }))

    }

    return processData
  } catch (err) {
    console.error('Error to load articles:', err)
  }
}
