const converter = new showdown.Converter({
  tables: true,
  simplifiedAutoLink: true,
  strikethrough: true,
  tasklists: true,
})
converter.setFlavor('github')

const content = document.getElementById('markdown')

const text = content.textContent
const html = converter.makeHtml(text)

content.remove()
document.getElementById('output').innerHTML = html
