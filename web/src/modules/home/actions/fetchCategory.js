export const fetchCategorie = async () => {
  try {
    const req = await fetch(`${API_URL}api/category`)
    if (req.status != 200) {
      throw new Error('Error to obtain data')
    }

    return req.json()
  } catch (err) {
    console.error(err)
  }
}
