export const FrontTags = ({ tags }) => {
  return (
    <div className='flex items-center flex-wrap gap-2 mt-2 w-full'>
      <span className='text-2xl mr-1 text-[#98c379]'></span>
      {tags?.map((tag, index) => (
        <span
          key={index}
          className='text-xs px-2 py-1 rounded-md bg-[#282c34] text-[#abb2bf] border border-[#3e4451]'
        >
          {tag}
        </span>
      ))}
    </div>
  )
}
