const formatDate = (date: Date) =>
  date.toLocaleString('es-PE', {
    year: 'numeric',
    month: 'long',
    day: '2-digit',
    timeZone: 'UTC'
  })

export default formatDate
