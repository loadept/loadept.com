
export const Footer = () => {
  const currentDate = new Date().getFullYear()

  return (
    <footer className="py-8 bg-linear-to-t from-[#282c34] via-80% via-[#282c34] to-[#1f2329]">
      <div className="container mx-auto px-4">
        <div className="flex flex-col items-center justify-center text-center">
          <p className="text-sm text-[#abb2bf]">
            <span className="text-[#c678dd]">año</span> <span class="text-[#e06c75]">:=</span> { currentDate }
          </p>
        </div>
      </div>
    </footer>
  )
}
