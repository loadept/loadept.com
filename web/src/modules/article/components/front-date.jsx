import formatDate from '../../../core/utils/format-date'

export const FrontDate = ({ date }) => {
  return (
    <div className='flex items-center mr-4'>
      <span className='text-2xl mr-2 text-[#c678dd]'></span>
      {formatDate(date, 'numeric')}
    </div>
  )
}
