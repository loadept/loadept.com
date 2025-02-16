const content = document.getElementById('markdown')

const text = content.textContent

const parsed = marked.parse(text)

content.remove()
document.getElementById('output').innerHTML = parsed 
