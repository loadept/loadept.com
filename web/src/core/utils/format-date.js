/**
  * @param {String} utcDate
*/
const formatDate = (utcDate, format) => {
  switch (format) {
    case "long":
      return new Date(utcDate).toLocaleString('es-PE', {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
      })
    case "numeric":
      return new Date(utcDate).toLocaleString('es-PE', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit'
      })
}
}

export default formatDate
