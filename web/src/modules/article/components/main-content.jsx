import { MarkdownContent } from './markdown-content'
import { FrontDate } from './front-date'
import { FrontTags } from './front-tags'
import { Spinner } from './spinner'
import { FeedBackCard } from './feed-back-card'

export const MainContent = ({ content, meta, repo_url }) => {
  return (
    <>
      {
        (meta.title && content) ?
          <>
            < article className='prose prose-invert prose-headings:text-[#e5c07b] prose-a:text-[#61afef] max-w-none' >
              <h1 className='text-5xl font-bold text-[#e5c07b] mb-4'>{meta.title.toUpperCase()}</h1>

              <div className='flex flex-wrap items-center text-sm text-[#5c6370] mb-8 gap-y-2'>
                <FrontDate date={meta.date} />
                <FrontTags tags={meta.tags} />
              </div>

              <MarkdownContent content={content} />
            </article >
            <FeedBackCard repo_url={repo_url} />
          </>
          :
          <div className='flex justify-center mt-48'>
            <Spinner />
          </div>
      }
    </>
  )
}
