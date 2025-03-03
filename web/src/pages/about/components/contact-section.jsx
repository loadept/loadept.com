export const ContactSection = () => {
  return (
    <section className="space-y-3 text-[#abb2bf]">
      <div className="flex items-center gap-3">
        <span className="text-3xl text-[#56b6c2]">󰺻</span>
        <code className="text-[#c678dd]">contacto.txt</code>
      </div>
      <div className="bg-[#282c34] rounded-lg p-6">
        <pre className="text-sm whitespace-pre-wrap">
          <code>{`# Para contactarme
email: hello@developer.com
github: github.com/developer
twitter: @developer

# O ejecuta este comando
npx contact-developer`}</code>
        </pre>
      </div>
    </section>
  )
}
