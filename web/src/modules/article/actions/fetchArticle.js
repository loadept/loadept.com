export const fetchArticle = async ({ category, name }) => {
  try {
    const req = await fetch(`${API_URL}api/article/${category}/${name}`)
    if (req.status != 200) {
      throw new Error('Error to obtain data')
    }

    return req.text()
  } catch (err) {
    console.error(err)
    return ''
  }
} 
