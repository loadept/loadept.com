import { useEffect, useState } from 'preact/hooks'

export const Categories = () => {
  const [data, setData] = useState({ categories: [] })

  console.log("API_URL =>", API_URL)

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const req = await fetch("http://localhost:8080/api/categories")
        const jsonData = await req.json()
        setData(jsonData)
      } catch (err) {
        console.log(err)
      }
    }

    fetchCategories()
  }, [])

  return (
    <section className="space-y-6 mt-15">
      <div className="flex items-center gap-3">
        <span className="text-3xl text-[#98c379]"></span>
        <h2 className="text-2xl font-bold text-[#e5c07b]">Contenido</h2>
      </div>

      <div className="space-y-6">
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-2 overflow-x-auto pb-2 scrollbar-hide">
            {data.categories.map((category, k) => (
              <p key={k}>{category.name}</p>
            ))}
          </div>
          <div className="relative flex items-center min-w-[200px]">
            <span className="text-3xl text-[#528bff] absolute left-3"></span>
            <input type="text" placeholder="Buscar..."
              className="w-full bg-[#282c34] rounded-md py-2 pl-10 pr-4 text-sm text-[#abb2bf] placeholder:text-[#528bff] focus:outline-none" />
          </div>
        </div>
      </div>
    </section>
  )
}
